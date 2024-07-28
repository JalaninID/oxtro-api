//go:build wireinject
// +build wireinject

package injector

import (
	"app/domain"
	"app/repository/repo_interest"
	"app/service/service_partner"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var newSetInterestRepository = wire.NewSet(
	repo_interest.NewRepository,
	wire.Bind(
		new(domain.RepositoryInterest),
		new(*repo_interest.Repository),
	),
)

func InitializedPartner(db *gorm.DB, logger *logrus.Logger) *service_partner.ServicePartner {
	wire.Build(newSetUserRepository, newSetInterestRepository, service_partner.NewService)
	return nil
}
