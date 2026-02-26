package middleware

import (
	"net/http"
	"strings"

	connectcors "connectrpc.com/cors"
)

// withCORS adds CORS support to a Connect HTTP handler.
func WithCORS(connectHandler http.Handler) http.Handler {
	allowedOrigins := map[string]struct{}{
		"http://localhost:5173": {},
		"http://localhost:5174": {},
		"http://127.0.0.1:5173": {},
		"http://127.0.0.1:5174": {},
		"http://localhost:4173": {},
		"http://127.0.0.1:4173": {},
	}
	allowedMethods := strings.Join(connectcors.AllowedMethods(), ",")
	allowedHeaders := strings.Join(append(connectcors.AllowedHeaders(), "Authorization"), ",")
	exposedHeaders := strings.Join(connectcors.ExposedHeaders(), ",")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if _, ok := allowedOrigins[origin]; ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Expose-Headers", exposedHeaders)
			w.Header().Set("Vary", "Origin")
		}

		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
				w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
				w.Header().Set("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}

		connectHandler.ServeHTTP(w, r)
	})
}
