package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
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

		statusText := fmt.Sprintf("%3d", recorder.status)
		latencyText := fmt.Sprintf("%10s", time.Since(start))
		ipText := fmt.Sprintf("%-15s", clientIP(r))
		methodText := fmt.Sprintf("%-7s", r.Method)

		logLine := fmt.Sprintf(
			"%s | %s | %s | %s | %s | %s | -",
			start.Local().Format("15:04:05"),
			colorStatus(statusText, recorder.status),
			latencyText,
			color.New(color.FgHiCyan).Sprint(ipText),
			colorMethod(methodText, r.Method),
			r.URL.Path,
		)
		logger.Debug(logLine)
	})
}

func colorStatus(statusText string, statusCode int) string {
	switch {
	case statusCode >= 500:
		return color.New(color.FgHiRed).Sprint(statusText)
	case statusCode >= 400:
		return color.New(color.FgHiYellow).Sprint(statusText)
	case statusCode >= 300:
		return color.New(color.FgHiBlue).Sprint(statusText)
	default:
		return color.New(color.FgHiGreen).Sprint(statusText)
	}
}

func colorMethod(methodText string, method string) string {
	switch method {
	case http.MethodGet:
		return color.New(color.FgHiGreen).Sprint(methodText)
	case http.MethodPost:
		return color.New(color.FgHiBlue).Sprint(methodText)
	case http.MethodPut, http.MethodPatch:
		return color.New(color.FgHiYellow).Sprint(methodText)
	case http.MethodDelete:
		return color.New(color.FgHiRed).Sprint(methodText)
	default:
		return color.New(color.FgHiWhite).Sprint(methodText)
	}
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
