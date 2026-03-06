package model

import "time"

type AppSetting struct {
	ID        int64
	Key       string `gorm:"size:191;uniqueIndex;not null"`
	Value     string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *AppSetting) TableName() string {
	return "app_settings"
}
