package routers

import (
	"app/gen/proto/user/v1/userv1connect"
	"app/injector"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) RouterUser() {
	_, handler := userv1connect.NewServiceUserHandler(
		injector.InitializedUser(r.config.Logger, r.config.Database),
	)
	r.Engine.POST(userv1connect.ServiceUserPurchasePremiumProcedure, middleware.Authorization(), gin.WrapH(handler))
}
