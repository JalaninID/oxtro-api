package routers

// import (
// 	"app/gen/proto/auth/v1/authv1connect"
// 	"app/injector"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func (r *Router) RouterAuth() {
// 	_, handler := authv1connect.NewServiceAuthHandler(
// 		injector.InitializedAuth(r.config.Database, r.config.Logger),
// 	)
// 	r.Engine.POST("/auth/*any", gin.WrapH(http.StripPrefix("/auth", handler)))
// }
