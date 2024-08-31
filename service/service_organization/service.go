package service_organization

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type service struct {
	serviceOrg domain.RepositoryOrganization
	logger     *logrus.Logger
}

func NewService(serviceOrg domain.RepositoryOrganization, logger *logrus.Logger) *service {
	return &service{serviceOrg: serviceOrg, logger: logger}
}
