package service_organization

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type service struct {
	repoOrg domain.RepositoryOrganization
	logger  *logrus.Logger
}

func NewService(repoOrg domain.RepositoryOrganization, logger *logrus.Logger) *service {
	return &service{repoOrg: repoOrg, logger: logger}
}
