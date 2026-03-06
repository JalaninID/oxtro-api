package service_setup

import (
	"app/constant"
	"app/dto/dto_setup"
	"app/model"
	"app/repository/repo_setup"
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestService(t *testing.T) *service {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.AppSetting{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}
	return NewService(repo_setup.NewRepository(db), db, logrus.New())
}

func TestGetStatus_DefaultIncomplete(t *testing.T) {
	svc := newTestService(t)

	status, err := svc.GetStatus(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if status.IsCompleted {
		t.Fatalf("expected setup to be incomplete")
	}
}

func TestRunSetup_SuccessAndIdempotent(t *testing.T) {
	t.Setenv("BCRYPT_COST", "4")
	svc := newTestService(t)

	_, err := svc.RunSetup(context.Background(), dto_setup.RunSetupRequest{
		Username: "admin",
		Email:    "admin@example.com",
		Password: "StrongPass1",
	})
	if err != nil {
		t.Fatalf("expected setup to succeed, got %v", err)
	}

	status, err := svc.GetStatus(context.Background())
	if err != nil {
		t.Fatalf("expected no error reading status, got %v", err)
	}
	if !status.IsCompleted {
		t.Fatalf("expected setup to be completed")
	}

	_, err = svc.RunSetup(context.Background(), dto_setup.RunSetupRequest{
		Username: "another-admin",
		Email:    "another@example.com",
		Password: "StrongPass1",
	})
	if !errors.Is(err, constant.ErrSetupAlreadyCompleted) {
		t.Fatalf("expected ErrSetupAlreadyCompleted, got %v", err)
	}
}
