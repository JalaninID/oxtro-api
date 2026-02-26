package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (w *statusRecorder) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func WithRequestLogger(next http.Handler, logger *logrus.Logger) http.Handler {
	if logger == nil {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		logger.Debugf(
			"%s | %d | %s | %s | %s | %s | -",
			start.Local().Format("15:04:05"),
			recorder.status,
			time.Since(start),
			clientIP(r),
			r.Method,
			r.URL.Path,
		)
	})
}

func clientIP(r *http.Request) string {
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		first := strings.Split(forwardedFor, ",")[0]
		trimmed := strings.TrimSpace(first)
		if trimmed != "" {
			return trimmed
		}
	}

	if realIP := strings.TrimSpace(r.Header.Get("X-Real-IP")); realIP != "" {
		return realIP
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}

	if r.RemoteAddr != "" {
		return r.RemoteAddr
	}

	return "-"
}
