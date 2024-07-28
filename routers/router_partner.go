package routers

import (
	"app/gen/proto/partner/v1/partnerv1connect"
	"app/injector"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) RouterPartner() {
	_, handler := partnerv1connect.NewServicePartnerHandler(
		injector.InitializedPartner(r.config.Database, r.config.Logger),
	)
	r.Engine.POST(partnerv1connect.ServicePartnerSearchPartnerProcedure, middleware.Authorization(), gin.WrapH(handler))
	r.Engine.POST(partnerv1connect.ServicePartnerSwipePartnerProcedure, middleware.Authorization(), gin.WrapH(handler))
}
