package service_partner

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type ServicePartner struct {
	repoPartner domain.RepositoryInterest
	repoUser    domain.RepositoryUser
	logger      *logrus.Logger
}

func NewService(logger *logrus.Logger, repoPartner domain.RepositoryInterest, repoUser domain.RepositoryUser) *ServicePartner {
	return &ServicePartner{
		repoPartner: repoPartner,
		logger:      logger,
		repoUser:    repoUser,
	}
}
