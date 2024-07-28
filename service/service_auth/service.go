package service_auth

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type ServiceAuth struct {
	logger   *logrus.Logger
	repoUser domain.RepositoryUser
}

func NewService(logger *logrus.Logger, repoUser domain.RepositoryUser) *ServiceAuth {
	return &ServiceAuth{
		logger:   logger,
		repoUser: repoUser,
	}
}
