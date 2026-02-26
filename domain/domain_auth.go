package domain

import (
	authv1 "app/gen/auth/v1"
	"app/model"
	"context"
)

//go:generate mockery --name RepositoryAuth
type RepositoryAuth interface {
	CreateSession(ctx context.Context, session model.AuthSession) (model.AuthSession, error)
	DetailSessionByTokenHash(ctx context.Context, tokenHash string) (model.AuthSession, error)
	DetailSessionByID(ctx context.Context, sessionID string, userID int) (model.AuthSession, error)
	ListSessions(ctx context.Context, userID int) ([]model.AuthSession, error)
	RevokeSessionByTokenHash(ctx context.Context, tokenHash string) error
	RevokeSessionByID(ctx context.Context, sessionID string, userID int) error
	RevokeAllSessionsByUserID(ctx context.Context, userID int) error
	RotateSession(ctx context.Context, oldSessionID int, session model.AuthSession) (model.AuthSession, error)
	CreatePasswordReset(ctx context.Context, reset model.PasswordReset) (model.PasswordReset, error)
	DetailPasswordReset(ctx context.Context, tokenHash string) (model.PasswordReset, error)
	MarkPasswordResetUsed(ctx context.Context, resetID int) error
	UpsertEmailVerification(ctx context.Context, verification model.EmailVerification) (model.EmailVerification, error)
	DetailEmailVerification(ctx context.Context, tokenHash string) (model.EmailVerification, error)
	MarkEmailVerificationUsed(ctx context.Context, verificationID int) error
	CreateLoginAttempt(ctx context.Context, attempt model.LoginAttempt) (model.LoginAttempt, error)
	CountRecentFailedLoginAttempts(ctx context.Context, userID int, sinceMinutes int) (int64, error)
	CreateSecurityEvent(ctx context.Context, event model.SecurityEvent) (model.SecurityEvent, error)
}

//go:generate mockery --name ServiceAuth
type ServiceAuth interface {
	Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error)
	Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error)
	RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error)
	Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.AuthStatusResponse, error)
	LogoutAll(ctx context.Context, req *authv1.LogoutAllRequest) (*authv1.AuthStatusResponse, error)
	ForgotPassword(ctx context.Context, req *authv1.ForgotPasswordRequest) (*authv1.ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, req *authv1.ResetPasswordRequest) (*authv1.AuthStatusResponse, error)
	ChangePassword(ctx context.Context, req *authv1.ChangePasswordRequest) (*authv1.AuthStatusResponse, error)
	VerifyEmail(ctx context.Context, req *authv1.VerifyEmailRequest) (*authv1.AuthStatusResponse, error)
	ResendEmailVerification(ctx context.Context, req *authv1.ResendEmailVerificationRequest) (*authv1.AuthStatusResponse, error)
	ListSessions(ctx context.Context, req *authv1.ListSessionsRequest) (*authv1.ListSessionsResponse, error)
	RevokeSession(ctx context.Context, req *authv1.RevokeSessionRequest) (*authv1.AuthStatusResponse, error)
}
