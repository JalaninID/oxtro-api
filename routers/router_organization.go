package routers

import (
	"app/gen/organization/v1/organizationv1connect"
	"app/injector"
	"app/middleware"
	"net/http"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
)

func (r *Router) RouterOrganization() {
	_, handler := organizationv1connect.NewOrganizationHandler(
		injector.InitializedOrganization(r.config.Database, r.config.Logger),
		connect.WithInterceptors(middleware.NewOrganizationInterceptor()),
	)

	r.Engine.POST("/organization/*any", gin.WrapH(http.StripPrefix("/organization", handler)))
}
