package service_auth

import (
	"app/constant"
	"app/domain"
	authv1 "app/gen/auth/v1"
	"app/middleware"
	"app/model"
	"app/pkg"
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type service struct {
	repoUser domain.RepositoryUser
	repoAuth domain.RepositoryAuth
	logger   *logrus.Logger
}

func NewService(repoUser domain.RepositoryUser, repoAuth domain.RepositoryAuth, logger *logrus.Logger) *service {
	return &service{
		repoUser: repoUser,
		repoAuth: repoAuth,
		logger:   logger,
	}
}

func (s *service) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" || req.GetUsername() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	if err := pkg.ValidatePassword(req.GetPassword()); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrPasswordPolicy)
	}

	existing, err := s.repoUser.DetailUser(ctx, model.FilterUser{Email: strings.ToLower(req.GetEmail())})
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Errorf("register detail user failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if existing.ID != 0 {
		return nil, connect.NewError(connect.CodeAlreadyExists, constant.ErrUsernameExist)
	}

	passwordHash, err := pkg.HashPassword(req.GetPassword())
	if err != nil {
		s.logger.Errorf("register hash password failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	user := model.User{
		UUID:     uuid.NewString(),
		Name:     req.GetUsername(),
		Email:    strings.ToLower(req.GetEmail()),
		Avatar:   req.GetProfile(),
		Bio:      req.GetName(),
		Password: passwordHash,
		IsActive: true,
	}
	createdUser, err := s.repoUser.CreateUser(ctx, user)
	if err != nil {
		s.logger.Errorf("register create user failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	if err := s.issueEmailVerification(ctx, createdUser.ID); err != nil {
		s.logger.Errorf("register issue email verification failed: %v", err)
	}

	return &authv1.RegisterResponse{
		Id:        createdUser.UUID,
		CreatedAt: createdUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdUser.UpdatedAt.Format(time.RFC3339),
		Username:  createdUser.Name,
		Email:     createdUser.Email,
		Name:      createdUser.Bio,
		Profile:   createdUser.Avatar,
		Status:    "pending_email_verification",
		Premium:   req.GetPremium(),
	}, nil
}

func (s *service) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}

	user, err := s.repoUser.DetailUser(ctx, model.FilterUser{Email: strings.ToLower(req.GetUsername())})
	if err == gorm.ErrRecordNotFound {
		_, _ = s.repoAuth.CreateLoginAttempt(ctx, model.LoginAttempt{
			UUID:      uuid.NewString(),
			Email:     strings.ToLower(req.GetUsername()),
			IPAddress: requestIPFromContext(ctx),
			Success:   false,
			Reason:    "user_not_found",
		})
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrInvalidCredentials)
	}
	if err != nil {
		s.logger.Errorf("login detail user failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, connect.NewError(connect.CodeResourceExhausted, constant.ErrUserLocked)
	}

	if err := pkg.ComparePassword(user.Password, req.GetPassword()); err != nil {
		_, _ = s.repoAuth.CreateLoginAttempt(ctx, model.LoginAttempt{
			UUID:      uuid.NewString(),
			UserID:    user.ID,
			Email:     user.Email,
			IPAddress: requestIPFromContext(ctx),
			Success:   false,
			Reason:    "password_mismatch",
		})
		_ = s.handlePotentialLock(ctx, user)
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrInvalidCredentials)
	}

	if !user.IsActive {
		return nil, connect.NewError(connect.CodePermissionDenied, constant.ErrAuthorization)
	}

	_, _ = s.repoAuth.CreateLoginAttempt(ctx, model.LoginAttempt{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		Email:     user.Email,
		IPAddress: requestIPFromContext(ctx),
		Success:   true,
		Reason:    "login_success",
	})

	tokenResponse, err := s.issueSessionTokens(ctx, user, req.GetDeviceId())
	if err != nil {
		s.logger.Errorf("login issue session tokens failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	if user.EmailVerifiedAt == nil {
		_, _ = s.repoAuth.CreateSecurityEvent(ctx, model.SecurityEvent{
			UUID:      uuid.NewString(),
			UserID:    user.ID,
			EventType: "login_unverified_email",
			Severity:  "warning",
			Metadata:  user.Email,
		})
	}

	return tokenResponse, nil
}

func (s *service) RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	if req.GetRefreshToken() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}

	tokenHash := pkg.HashToken(req.GetRefreshToken())
	session, err := s.repoAuth.DetailSessionByTokenHash(ctx, tokenHash)
	if err == gorm.ErrRecordNotFound {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if err != nil {
		s.logger.Errorf("refresh detail session failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	if session.RevokedAt != nil || session.ExpiresAt.Before(time.Now()) {
		_, _ = s.repoAuth.CreateSecurityEvent(ctx, model.SecurityEvent{
			UUID:      uuid.NewString(),
			UserID:    session.UserID,
			EventType: "refresh_reuse_or_expired",
			Severity:  "high",
			Metadata:  session.DeviceID,
		})
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if req.GetDeviceId() != "" && session.DeviceID != "" && req.GetDeviceId() != session.DeviceID {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}

	user, err := s.repoUser.DetailUser(ctx, model.FilterUser{ID: session.UserID})
	if err != nil {
		s.logger.Errorf("refresh detail user failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	refreshTTL := refreshTokenTTL()
	newRefreshToken := pkg.GenerateRandomString(96)
	newSession := model.AuthSession{
		UUID:          uuid.NewString(),
		UserID:        user.ID,
		TokenHash:     pkg.HashToken(newRefreshToken),
		DeviceID:      session.DeviceID,
		UserAgent:     requestUserAgentFromContext(ctx),
		IPAddress:     requestIPFromContext(ctx),
		ExpiresAt:     time.Now().Add(refreshTTL),
		RotatedFromID: &session.ID,
	}

	rotated, err := s.repoAuth.RotateSession(ctx, session.ID, newSession)
	if err != nil {
		s.logger.Errorf("refresh rotate session failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	accessTTL := accessTokenTTL()
	accessToken, err := pkg.Sign(map[string]any{
		"id":         user.UUID,
		"session_id": rotated.UUID,
		"token_type": "access",
	}, int(accessTTL.Minutes()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return &authv1.RefreshTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  time.Now().Add(accessTTL).Unix(),
		RefreshTokenExpiresAt: rotated.ExpiresAt.Unix(),
	}, nil
}

func (s *service) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetRefreshToken() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	if err := s.repoAuth.RevokeSessionByTokenHash(ctx, pkg.HashToken(req.GetRefreshToken())); err != nil {
		s.logger.Errorf("logout revoke session failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return &authv1.AuthStatusResponse{Success: true, Message: "logged out"}, nil
}

func (s *service) LogoutAll(ctx context.Context, _ *authv1.LogoutAllRequest) (*authv1.AuthStatusResponse, error) {
	user, err := s.currentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.repoAuth.RevokeAllSessionsByUserID(ctx, user.ID); err != nil {
		s.logger.Errorf("logout all revoke sessions failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return &authv1.AuthStatusResponse{Success: true, Message: "all sessions revoked"}, nil
}

func (s *service) ForgotPassword(ctx context.Context, req *authv1.ForgotPasswordRequest) (*authv1.ForgotPasswordResponse, error) {
	if req.GetEmail() == "" {
		return &authv1.ForgotPasswordResponse{Accepted: true}, nil
	}

	user, err := s.repoUser.DetailUser(ctx, model.FilterUser{Email: strings.ToLower(req.GetEmail())})
	if err == gorm.ErrRecordNotFound {
		return &authv1.ForgotPasswordResponse{Accepted: true}, nil
	}
	if err != nil {
		s.logger.Errorf("forgot password detail user failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	token := pkg.GenerateRandomString(80)
	_, err = s.repoAuth.CreatePasswordReset(ctx, model.PasswordReset{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		TokenHash: pkg.HashToken(token),
		ExpiresAt: time.Now().Add(30 * time.Minute),
	})
	if err != nil {
		s.logger.Errorf("forgot password create reset failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	s.logger.Infof("password reset token issued for %s: %s", user.Email, token)
	return &authv1.ForgotPasswordResponse{Accepted: true}, nil
}

func (s *service) ResetPassword(ctx context.Context, req *authv1.ResetPasswordRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetToken() == "" || req.GetNewPassword() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	if err := pkg.ValidatePassword(req.GetNewPassword()); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrPasswordPolicy)
	}

	reset, err := s.repoAuth.DetailPasswordReset(ctx, pkg.HashToken(req.GetToken()))
	if err == gorm.ErrRecordNotFound || (reset.UsedAt != nil) || reset.ExpiresAt.Before(time.Now()) {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if err != nil {
		s.logger.Errorf("reset password detail reset failed: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	passwordHash, err := pkg.HashPassword(req.GetNewPassword())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if _, err := s.repoUser.UpdateUser(ctx, model.FilterUser{ID: reset.UserID}, model.User{
		Password: passwordHash,
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if err := s.repoAuth.MarkPasswordResetUsed(ctx, reset.ID); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	_ = s.repoAuth.RevokeAllSessionsByUserID(ctx, reset.UserID)

	return &authv1.AuthStatusResponse{Success: true, Message: "password reset successful"}, nil
}

func (s *service) ChangePassword(ctx context.Context, req *authv1.ChangePasswordRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetOldPassword() == "" || req.GetNewPassword() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	if err := pkg.ValidatePassword(req.GetNewPassword()); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrPasswordPolicy)
	}

	user, err := s.currentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := pkg.ComparePassword(user.Password, req.GetOldPassword()); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrPassword)
	}
	if err := pkg.ComparePassword(user.Password, req.GetNewPassword()); err == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrPasswordPolicy)
	}

	passwordHash, err := pkg.HashPassword(req.GetNewPassword())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if _, err := s.repoUser.UpdateUser(ctx, model.FilterUser{ID: user.ID}, model.User{
		Password: passwordHash,
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	_ = s.repoAuth.RevokeAllSessionsByUserID(ctx, user.ID)
	return &authv1.AuthStatusResponse{Success: true, Message: "password changed"}, nil
}

func (s *service) VerifyEmail(ctx context.Context, req *authv1.VerifyEmailRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetToken() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	verification, err := s.repoAuth.DetailEmailVerification(ctx, pkg.HashToken(req.GetToken()))
	if err == gorm.ErrRecordNotFound || verification.UsedAt != nil || verification.ExpiresAt.Before(time.Now()) {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	now := time.Now()
	if _, err := s.repoUser.UpdateUser(ctx, model.FilterUser{ID: verification.UserID}, model.User{
		EmailVerifiedAt: &now,
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if err := s.repoAuth.MarkEmailVerificationUsed(ctx, verification.ID); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return &authv1.AuthStatusResponse{Success: true, Message: "email verified"}, nil
}

func (s *service) ResendEmailVerification(ctx context.Context, req *authv1.ResendEmailVerificationRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetEmail() == "" {
		return &authv1.AuthStatusResponse{Success: true, Message: "accepted"}, nil
	}
	user, err := s.repoUser.DetailUser(ctx, model.FilterUser{Email: strings.ToLower(req.GetEmail())})
	if err == gorm.ErrRecordNotFound {
		return &authv1.AuthStatusResponse{Success: true, Message: "accepted"}, nil
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if err := s.issueEmailVerification(ctx, user.ID); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return &authv1.AuthStatusResponse{Success: true, Message: "verification sent"}, nil
}

func (s *service) ListSessions(ctx context.Context, _ *authv1.ListSessionsRequest) (*authv1.ListSessionsResponse, error) {
	user, err := s.currentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	sessions, err := s.repoAuth.ListSessions(ctx, user.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	response := &authv1.ListSessionsResponse{}
	for _, item := range sessions {
		response.Sessions = append(response.Sessions, &authv1.Session{
			Id:        item.UUID,
			DeviceId:  item.DeviceID,
			UserAgent: item.UserAgent,
			Ip:        item.IPAddress,
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			ExpiresAt: item.ExpiresAt.Format(time.RFC3339),
			Revoked:   item.RevokedAt != nil,
		})
	}
	return response, nil
}

func (s *service) RevokeSession(ctx context.Context, req *authv1.RevokeSessionRequest) (*authv1.AuthStatusResponse, error) {
	if req.GetSessionId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidCredentials)
	}
	user, err := s.currentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.repoAuth.RevokeSessionByID(ctx, req.GetSessionId(), user.ID); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return &authv1.AuthStatusResponse{Success: true, Message: "session revoked"}, nil
}

func (s *service) currentUserFromContext(ctx context.Context) (model.User, error) {
	principal, ok := middleware.PrincipalFromContext(ctx)
	if !ok || principal.Token.ID == "" {
		return model.User{}, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	user, err := s.repoUser.DetailUser(ctx, model.FilterUser{UUID: principal.Token.ID})
	if err == gorm.ErrRecordNotFound {
		return model.User{}, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if err != nil {
		s.logger.Errorf("current user lookup failed: %v", err)
		return model.User{}, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return user, nil
}

func (s *service) issueSessionTokens(ctx context.Context, user model.User, deviceID string) (*authv1.LoginResponse, error) {
	if deviceID == "" {
		deviceID = "unknown-device"
	}

	refreshTTL := refreshTokenTTL()
	accessTTL := accessTokenTTL()
	refreshToken := pkg.GenerateRandomString(96)
	session := model.AuthSession{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		TokenHash: pkg.HashToken(refreshToken),
		DeviceID:  deviceID,
		UserAgent: requestUserAgentFromContext(ctx),
		IPAddress: requestIPFromContext(ctx),
		ExpiresAt: time.Now().Add(refreshTTL),
	}
	created, err := s.repoAuth.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	accessToken, err := pkg.Sign(map[string]any{
		"id":         user.UUID,
		"session_id": created.UUID,
		"token_type": "access",
	}, int(accessTTL.Minutes()))
	if err != nil {
		return nil, err
	}

	return &authv1.LoginResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  time.Now().Add(accessTTL).Unix(),
		RefreshTokenExpiresAt: created.ExpiresAt.Unix(),
	}, nil
}

func (s *service) issueEmailVerification(ctx context.Context, userID int) error {
	token := pkg.GenerateRandomString(80)
	_, err := s.repoAuth.UpsertEmailVerification(ctx, model.EmailVerification{
		UUID:      uuid.NewString(),
		UserID:    userID,
		TokenHash: pkg.HashToken(token),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		return err
	}
	s.logger.Infof("email verification token issued for user_id=%d: %s", userID, token)
	return nil
}

func (s *service) handlePotentialLock(ctx context.Context, user model.User) error {
	const failedAttemptThreshold = 5
	const failedAttemptWindowMinutes = 15
	const lockDurationMinutes = 15

	failedCount, err := s.repoAuth.CountRecentFailedLoginAttempts(ctx, user.ID, failedAttemptWindowMinutes)
	if err != nil {
		return err
	}
	if failedCount < failedAttemptThreshold {
		return nil
	}
	lockedUntil := time.Now().Add(lockDurationMinutes * time.Minute)
	_, err = s.repoUser.UpdateUser(ctx, model.FilterUser{ID: user.ID}, model.User{
		LockedUntil: &lockedUntil,
	})
	if err != nil {
		return err
	}
	_, _ = s.repoAuth.CreateSecurityEvent(ctx, model.SecurityEvent{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		EventType: "account_locked",
		Severity:  "medium",
		Metadata:  lockedUntil.Format(time.RFC3339),
	})
	return nil
}

func accessTokenTTL() time.Duration {
	minutes := 15
	if parsed, err := strconv.Atoi(getEnv("AUTH_ACCESS_TOKEN_MINUTES", "15")); err == nil && parsed > 0 {
		minutes = parsed
	}
	return time.Duration(minutes) * time.Minute
}

func refreshTokenTTL() time.Duration {
	hours := 24 * 7
	if parsed, err := strconv.Atoi(getEnv("AUTH_REFRESH_TOKEN_HOURS", "168")); err == nil && parsed > 0 {
		hours = parsed
	}
	return time.Duration(hours) * time.Hour
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
