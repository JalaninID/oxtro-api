package injector

import (
	"app/domain"
	"app/handler/handler_auth"
	"app/repository/repo_auth"
	"app/repository/repo_user"
	"app/service/service_auth"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewSetRepositoryUser(db *gorm.DB) domain.RepositoryUser {
	return repo_user.NewRepository(db)
}

func NewSetRepositoryAuth(db *gorm.DB) domain.RepositoryAuth {
	return repo_auth.NewRepository(db)
}

func NewSetServiceAuth(repoUser domain.RepositoryUser, repoAuth domain.RepositoryAuth, logger *logrus.Logger) domain.ServiceAuth {
	return service_auth.NewService(repoUser, repoAuth, logger)
}

func InitializedAuth(db *gorm.DB, logger *logrus.Logger) *handler_auth.Auth {
	repoUser := NewSetRepositoryUser(db)
	repoAuth := NewSetRepositoryAuth(db)
	service := NewSetServiceAuth(repoUser, repoAuth, logger)
	handler := handler_auth.NewHandlerAuth(service)
	return handler
}
