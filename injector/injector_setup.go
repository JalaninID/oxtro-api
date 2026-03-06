package injector

import (
	"app/domain"
	"app/handler/handler_setup"
	"app/repository/repo_setup"
	"app/service/service_setup"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewSetRepositorySetup(db *gorm.DB) domain.RepositorySetup {
	return repo_setup.NewRepository(db)
}

func NewSetServiceSetup(repoSetup domain.RepositorySetup, db *gorm.DB, logger *logrus.Logger) domain.ServiceSetup {
	return service_setup.NewService(repoSetup, db, logger)
}

func InitializedSetup(db *gorm.DB, logger *logrus.Logger) *handler_setup.Setup {
	repoSetup := NewSetRepositorySetup(db)
	service := NewSetServiceSetup(repoSetup, db, logger)
	handler := handler_setup.NewHandlerSetup(service)
	return handler
}
