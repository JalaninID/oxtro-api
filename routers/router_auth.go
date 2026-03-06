package routers

import (
	"app/gen/auth/v1/authv1connect"
	"app/injector"
	"app/middleware"
	"context"

	"connectrpc.com/connect"
)

func (r *Router) RouterAuth() {
	publicProcedures := map[string]struct{}{
		authv1connect.AuthLoginProcedure:                   {},
		authv1connect.AuthRefreshTokenProcedure:            {},
		authv1connect.AuthForgotPasswordProcedure:          {},
		authv1connect.AuthResetPasswordProcedure:           {},
		authv1connect.AuthVerifyEmailProcedure:             {},
		authv1connect.AuthResendEmailVerificationProcedure: {},
	}
	isCompleted, err := r.setupService.IsSetupCompleted(context.Background())
	if err != nil {
		r.config.Logger.Warnf("setup status check failed on auth router init: %v", err)
	}
	if isCompleted {
		publicProcedures[authv1connect.AuthRegisterProcedure] = struct{}{}
	}

	path, handler := authv1connect.NewAuthHandler(
		injector.InitializedAuth(r.config.Database, r.config.Logger, r.hooks),
		connect.WithInterceptors(middleware.NewAuthInterceptor(publicProcedures)),
	)
	r.Mux.Handle(path, handler)
}
