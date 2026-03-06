package routers

import (
	"app/gen/setup/v1/setupv1connect"
	"app/injector"
	"app/middleware"

	"connectrpc.com/connect"
)

func (r *Router) RouterSetup() {
	publicProcedures := map[string]struct{}{
		setupv1connect.SetupGetStatusProcedure: {},
		setupv1connect.SetupRunSetupProcedure:  {},
	}
	path, handler := setupv1connect.NewSetupHandler(
		injector.InitializedSetup(r.config.Database, r.config.Logger),
		connect.WithInterceptors(middleware.NewAuthInterceptor(publicProcedures)),
	)
	r.Mux.Handle(path, handler)
}
