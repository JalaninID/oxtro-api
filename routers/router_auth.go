package routers

import (
	"app/gen/proto/auth/v1/authv1connect"
	"app/injector"

	"github.com/gin-gonic/gin"
)

func (r *Router) RouterAuth() {
	_, handler := authv1connect.NewServiceAuthHandler(
		injector.InitializedAuth(r.config.Database, r.config.Logger),
	)
	r.Engine.POST(authv1connect.ServiceAuthLoginProcedure, gin.WrapH(handler))
	r.Engine.POST(authv1connect.ServiceAuthRegisterProcedure, gin.WrapH(handler))
}
