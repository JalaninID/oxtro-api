//go:build wireinject
// +build wireinject

package injector

import (
	"app/crons/cron_interest"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitializedCronInterest(logger *logrus.Logger, db *gorm.DB) *cron_interest.CronInterest {
	wire.Build(newSetInterestRepository, cron_interest.NewCronInterest)
	return nil
}
