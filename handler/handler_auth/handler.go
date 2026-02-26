package handler_auth

import (
	"app/domain"
	authv1 "app/gen/auth/v1"
	"app/gen/auth/v1/authv1connect"
	"app/service/service_auth"
	"context"
	"strings"

	"connectrpc.com/connect"
)

type Auth struct {
	serviceAuth domain.ServiceAuth
	authv1connect.UnimplementedAuthHandler
}

func NewHandlerAuth(serviceAuth domain.ServiceAuth) *Auth {
	return &Auth{serviceAuth: serviceAuth}
}

func (h *Auth) Register(ctx context.Context, req *connect.Request[authv1.RegisterRequest]) (*connect.Response[authv1.RegisterResponse], error) {
	response, err := h.serviceAuth.Register(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	response, err := h.serviceAuth.Login(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) RefreshToken(ctx context.Context, req *connect.Request[authv1.RefreshTokenRequest]) (*connect.Response[authv1.RefreshTokenResponse], error) {
	response, err := h.serviceAuth.RefreshToken(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) Logout(ctx context.Context, req *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.Logout(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) LogoutAll(ctx context.Context, req *connect.Request[authv1.LogoutAllRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.LogoutAll(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) ForgotPassword(ctx context.Context, req *connect.Request[authv1.ForgotPasswordRequest]) (*connect.Response[authv1.ForgotPasswordResponse], error) {
	response, err := h.serviceAuth.ForgotPassword(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) ResetPassword(ctx context.Context, req *connect.Request[authv1.ResetPasswordRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.ResetPassword(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) ChangePassword(ctx context.Context, req *connect.Request[authv1.ChangePasswordRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.ChangePassword(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) VerifyEmail(ctx context.Context, req *connect.Request[authv1.VerifyEmailRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.VerifyEmail(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) ResendEmailVerification(ctx context.Context, req *connect.Request[authv1.ResendEmailVerificationRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.ResendEmailVerification(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) ListSessions(ctx context.Context, req *connect.Request[authv1.ListSessionsRequest]) (*connect.Response[authv1.ListSessionsResponse], error) {
	response, err := h.serviceAuth.ListSessions(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Auth) RevokeSession(ctx context.Context, req *connect.Request[authv1.RevokeSessionRequest]) (*connect.Response[authv1.AuthStatusResponse], error) {
	response, err := h.serviceAuth.RevokeSession(withRequestMetadata(ctx, req), req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func withRequestMetadata[T any](ctx context.Context, req *connect.Request[T]) context.Context {
	userAgent := req.Header().Get("User-Agent")
	ip := req.Header().Get("X-Forwarded-For")
	if ip == "" {
		ip = req.Header().Get("X-Real-IP")
	}
	if ip == "" {
		ip = "unknown"
	}
	ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	return service_auth.WithRequestMetadataForHandler(ctx, ip, userAgent)
}
