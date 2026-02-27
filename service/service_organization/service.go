package service_organization

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type service struct {
	repoOrg domain.RepositoryOrganization
	logger  *logrus.Logger
	hooks   domain.HookDispatcher
}

func NewService(repoOrg domain.RepositoryOrganization, logger *logrus.Logger, hooks domain.HookDispatcher) *service {
	return &service{repoOrg: repoOrg, logger: logger, hooks: hooks}
}
