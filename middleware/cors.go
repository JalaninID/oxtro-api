package middleware

import (
	"net/http"

	connectcors "connectrpc.com/cors"

	"github.com/rs/cors"
)

// withCORS adds CORS support to a Connect HTTP handler.
func WithCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return c.Handler(connectHandler)
}
