//go:build wireinject
// +build wireinject

package injector

import (
	"app/domain"
	"app/repository/repo_user"
	"app/service/service_user"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var newSetUserRepository = wire.NewSet(
	repo_user.NewRepository,
	wire.Bind(
		new(domain.RepositoryUser),
		new(*repo_user.Repository),
	),
)

// InitializedUser is a function to initialize user service
func InitializedUser(logger *logrus.Logger, db *gorm.DB) *service_user.ServiceUser {
	wire.Build(newSetUserRepository, service_user.NewService)
	return nil
}
