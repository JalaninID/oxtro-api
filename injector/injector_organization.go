package injector

import (
	"app/domain"
	"app/handler/handler_organization"
	"app/repository/repo_organization"
	"app/service/service_organization"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewSetRepositoryOrganization(db *gorm.DB) domain.RepositoryOrganization {
	return repo_organization.NewRepository(db)
}
func NewSetServiceOrganization(repo domain.RepositoryOrganization, logger *logrus.Logger, hooks domain.HookDispatcher) domain.ServiceOrganization {
	return service_organization.NewService(repo, logger, hooks)
}

func InitializedOrganization(db *gorm.DB, logger *logrus.Logger, hooks domain.HookDispatcher) *handler_organization.Organization {
	repo := NewSetRepositoryOrganization(db)
	service := NewSetServiceOrganization(repo, logger, hooks)
	handler := handler_organization.NewHandlerOrganization(service)
	return handler
}
