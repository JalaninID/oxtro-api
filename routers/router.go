package routers

import (
	"app/config"
	"app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Router struct {
	config *config.Config
	Engine *gin.Engine
}

func NewRouter(config *config.Config) *Router {
	engine := gin.Default()
	return &Router{
		config: config,
		Engine: engine,
	}
}
func (r *Router) Run() error {
	handler := middleware.WithCORS(r.Engine)
	return http.ListenAndServe(":8080", h2c.NewHandler(handler, &http2.Server{}))
}
