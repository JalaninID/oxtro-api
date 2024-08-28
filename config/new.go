package config

import (
	"app/pkg"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Config is a struct to store the config
type Config struct {
	// Database is a pointer to the database
	Database *gorm.DB
	// Logger is a pointer to the logger
	Logger *logrus.Logger
}

// NewConfig is a function to create a new config
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
