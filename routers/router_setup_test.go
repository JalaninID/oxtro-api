package routers

import (
	"app/config"
	"app/gen/setup/v1/setupv1connect"
	"app/middleware"
	"app/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSetupRouterTestHandler(t *testing.T) http.Handler {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.AppSetting{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	router := NewRouter(&config.Config{
		Database: db,
		Logger:   logrus.New(),
	}, nil)
	router.RouterSetup()

	return middleware.WithSetupGate(router.Mux, router.setupService)
}

func TestSetupWizardSmokeFlow(t *testing.T) {
	t.Setenv("BCRYPT_COST", "4")
	handler := newSetupRouterTestHandler(t)

	// Before setup, status should be incomplete.
	{
		reqBody := "{}"
		req := httptest.NewRequest(http.MethodPost, setupv1connect.SetupGetStatusProcedure, strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connect-Protocol-Version", "1")
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("expected status status code 200, got %d", rr.Code)
		}
	}

	// Bootstrap setup over Connect endpoint.
	{
		req := httptest.NewRequest(http.MethodPost, setupv1connect.SetupRunSetupProcedure, strings.NewReader(`{"username":"admin","email":"admin@example.com","password":"StrongPass1"}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connect-Protocol-Version", "1")
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("expected bootstrap status 200, got %d", rr.Code)
		}
	}

	// Setup cannot run twice.
	{
		req := httptest.NewRequest(http.MethodPost, setupv1connect.SetupRunSetupProcedure, strings.NewReader(`{"username":"admin","email":"admin@example.com","password":"StrongPass1"}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connect-Protocol-Version", "1")
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusConflict {
			t.Fatalf("expected second bootstrap status 409, got %d", rr.Code)
		}
	}

	// After setup, setup run endpoint should stay locked.
	{
		req := httptest.NewRequest(http.MethodPost, setupv1connect.SetupRunSetupProcedure, strings.NewReader(`{"username":"admin","email":"admin@example.com","password":"StrongPass1"}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connect-Protocol-Version", "1")
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusConflict {
			t.Fatalf("expected setup endpoint locked with 409, got %d", rr.Code)
		}
	}

}
