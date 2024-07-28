package config

import (
	"app/pkg"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Config struct {
	Database *gorm.DB
	Logger   *logrus.Logger
}

func NewConfig() *Config {
	db, err := NewPostgre()
	if err != nil {
		panic(err)
	}

	return &Config{
		Database: db,
		Logger:   pkg.NewLogger(),
	}
}
