package middleware

import (
	"app/gen/setup/v1/setupv1connect"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeSetupChecker struct {
	completed bool
	err       error
}

func (f fakeSetupChecker) IsSetupCompleted(_ context.Context) (bool, error) {
	return f.completed, f.err
}

func TestWithSetupGate_BlocksNonSetupPathWhenIncomplete(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := WithSetupGate(next, fakeSetupChecker{completed: false})

	req := httptest.NewRequest(http.MethodGet, "/auth.v1.Auth/ListSessions", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", rr.Code)
	}
}

func TestWithSetupGate_AllowsSetupStatusWhenIncomplete(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := WithSetupGate(next, fakeSetupChecker{completed: false})

	req := httptest.NewRequest(http.MethodGet, setupv1connect.SetupGetStatusProcedure, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestWithSetupGate_BlocksSetupPageWhenCompleted(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := WithSetupGate(next, fakeSetupChecker{completed: true})

	req := httptest.NewRequest(http.MethodGet, setupv1connect.SetupRunSetupProcedure, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusConflict {
		t.Fatalf("expected status 409, got %d", rr.Code)
	}
}

func TestWithSetupGate_ReportsCheckerError(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := WithSetupGate(next, fakeSetupChecker{err: errors.New("db failure")})

	req := httptest.NewRequest(http.MethodGet, setupv1connect.SetupGetStatusProcedure, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
}
