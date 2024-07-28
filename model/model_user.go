package model

import (
	"time"
)

const (
	statusActive = "active"
	statusBanned = "banned"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Username  string
	Password  string
	Name      string
	Profile   string
	Status    string
	IsPremium bool
	Gender    string
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) StatusActive() {
	u.Status = statusActive
}

func (u *User) StatusBanned() {
	u.Status = statusBanned
}
