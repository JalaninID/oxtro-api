package domain

import (
	"app/dto/dto_setup"
	"app/model"
	"context"
)

//go:generate mockery --name RepositorySetup
type RepositorySetup interface {
	DetailAppSetting(ctx context.Context, key string) (model.AppSetting, error)
	CreateAppSetting(ctx context.Context, setting model.AppSetting) (model.AppSetting, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

//go:generate mockery --name ServiceSetup
type ServiceSetup interface {
	GetStatus(ctx context.Context) (*dto_setup.SetupStatusResponse, error)
	RunSetup(ctx context.Context, req dto_setup.RunSetupRequest) (*dto_setup.RunSetupResponse, error)
	IsSetupCompleted(ctx context.Context) (bool, error)
}
