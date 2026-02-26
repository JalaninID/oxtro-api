package routers

import (
	"app/config"
	"app/middleware"
	"app/pkg"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Router struct {
	config *config.Config
	Mux    *http.ServeMux
}

func startupBanner(addr string) {
	fmt.Printf(`
 ____ ___  _ _____ ____  ____ 
/  _ \\  \///__ __Y  __\/  _ \
| / \| \  /   / \ |  \/|| / \|
| \_/| /  \   | | |    /| \_/|
\____//__/\\  \_/ \_/\_\\____/
                              

OXTRO API is starting...
Listening on %s
`, addr)
	fmt.Println()
}

func NewRouter(config *config.Config) *Router {
	mux := http.NewServeMux()
	return &Router{
		config: config,
		Mux:    mux,
	}
}
func (r *Router) Run() error {
	handler := middleware.WithRequestLogger(middleware.WithCORS(r.Mux), pkg.NewRequestDebugLogger())
	server := &http.Server{
		Addr:              ":8080",
		Handler:           h2c.NewHandler(handler, &http2.Server{CountError: func(errType string) { fmt.Println(errType) }}),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	startupBanner(server.Addr)
	return server.ListenAndServe()
}
