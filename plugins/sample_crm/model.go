package sample_crm

import "time"

type CRMContact struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	UUID      string     `gorm:"uniqueIndex;size:36;not null"`
	Name      string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:255"`
	Phone     string     `gorm:"size:50"`
	Company   string     `gorm:"size:255"`
	Notes     string     `gorm:"type:text"`
	CreatedAt time.Time  `gorm:"not null;default:now()"`
	UpdatedAt time.Time  `gorm:"not null;default:now()"`
	DeletedAt *time.Time `gorm:"index"`
}

func (CRMContact) TableName() string {
	return "plg_crm_contacts"
}
