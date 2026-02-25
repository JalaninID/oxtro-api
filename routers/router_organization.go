package routers

import (
	"app/gen/organization/v1/organizationv1connect"
	"app/injector"
	"app/middleware"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
)

func (r *Router) RouterOrganization() {
	path, handler := organizationv1connect.NewOrganizationHandler(
		injector.InitializedOrganization(r.config.Database, r.config.Logger),
		connect.WithInterceptors(middleware.NewAuthInterceptor(nil)),
	)

	r.Engine.Any(path+"*any", gin.WrapH(handler))
}
