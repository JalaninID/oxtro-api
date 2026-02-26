package routers

import (
	"app/gen/auth/v1/authv1connect"
	"app/injector"
	"app/middleware"

	"connectrpc.com/connect"
)

func (r *Router) RouterAuth() {
	publicProcedures := map[string]struct{}{
		authv1connect.AuthRegisterProcedure:                {},
		authv1connect.AuthLoginProcedure:                   {},
		authv1connect.AuthRefreshTokenProcedure:            {},
		authv1connect.AuthForgotPasswordProcedure:          {},
		authv1connect.AuthResetPasswordProcedure:           {},
		authv1connect.AuthVerifyEmailProcedure:             {},
		authv1connect.AuthResendEmailVerificationProcedure: {},
	}

	path, handler := authv1connect.NewAuthHandler(
		injector.InitializedAuth(r.config.Database, r.config.Logger),
		connect.WithInterceptors(middleware.NewAuthInterceptor(publicProcedures)),
	)
	r.Mux.Handle(path, handler)
}
