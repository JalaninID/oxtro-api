//go:build wireinject
// +build wireinject

package injector

import (
	"app/service/service_auth"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitializedAuth(db *gorm.DB, logger *logrus.Logger) *service_auth.ServiceAuth {
	wire.Build(newSetUserRepository, service_auth.NewService)
	return nil
}
