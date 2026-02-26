package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWithCORS_PreflightAllowsConnectAndAuthorizationHeaders(t *testing.T) {
	wrapped := WithCORS(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodOptions, "/organization.v1.Organization/ListOrganization", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Access-Control-Request-Method", http.MethodPost)
	req.Header.Set("Access-Control-Request-Headers", "Content-Type,Connect-Protocol-Version,Authorization")

	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}

	origin := req.Header.Get("Origin")
	allowOrigin := rec.Header().Get("Access-Control-Allow-Origin")
	if allowOrigin != origin {
		t.Fatalf("expected Access-Control-Allow-Origin to be %q, got %q", origin, allowOrigin)
	}

	allowMethods := rec.Header().Get("Access-Control-Allow-Methods")
	if !strings.Contains(allowMethods, http.MethodPost) {
		t.Fatalf("expected POST in Access-Control-Allow-Methods, got %q", allowMethods)
	}

	allowHeaders := rec.Header().Get("Access-Control-Allow-Headers")
	if !strings.Contains(strings.ToLower(allowHeaders), "authorization") {
		t.Fatalf("expected Authorization in Access-Control-Allow-Headers, got %q", allowHeaders)
	}
}
