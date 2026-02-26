package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestWithRequestLogger_LogsFiberLikeFields(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := logrus.New()
	logger.SetOutput(buffer)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		DisableColors:    true,
	})

	handler := WithRequestLogger(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}), logger)

	req := httptest.NewRequest(http.MethodPost, "/organization.v1.Organization/ListOrganization", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	logLine := buffer.String()
	expectedParts := []string{
		" | 201 | ",
		" | 127.0.0.1 | ",
		" | POST | ",
		" | /organization.v1.Organization/ListOrganization | -",
	}
	for _, part := range expectedParts {
		if !strings.Contains(logLine, part) {
			t.Fatalf("expected log to contain %q, got %q", part, logLine)
		}
	}
}

func TestClientIP_PrioritizesForwardedHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.7, 70.41.3.18")
	req.Header.Set("X-Real-IP", "198.51.100.9")
	req.RemoteAddr = "127.0.0.1:9000"

	got := clientIP(req)
	if got != "203.0.113.7" {
		t.Fatalf("expected first forwarded IP, got %q", got)
	}
}
