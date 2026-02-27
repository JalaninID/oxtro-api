package routers

import (
	"app/gen/organization/v1/organizationv1connect"
	"app/injector"
	"app/middleware"

	"connectrpc.com/connect"
)

func (r *Router) RouterOrganization() {
	path, handler := organizationv1connect.NewOrganizationHandler(
		injector.InitializedOrganization(r.config.Database, r.config.Logger, r.hooks),
		connect.WithInterceptors(middleware.NewAuthInterceptor(nil)),
	)

	r.Mux.Handle(path, handler)
}
