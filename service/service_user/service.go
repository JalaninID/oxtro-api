package service_user

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type ServiceUser struct {
	logger *logrus.Logger
	repo   domain.RepositoryUser
}

func NewService(logger *logrus.Logger, repo domain.RepositoryUser) *ServiceUser {
	return &ServiceUser{
		logger: logger,
		repo:   repo,
	}
}
