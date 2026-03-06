package model

import (
	"time"
)

// User represents a user
type (
	User struct {
		// ID is the primary key
		ID int
		// CreatedAt is the time the user was created
		CreatedAt time.Time
		// UpdatedAt is the time the user was last updated
		UpdatedAt time.Time
		// DeletedAt is the time the user was deleted
		DeletedAt *time.Time
		// UUID is the unique identifier for the user
		UUID string
		// Name is the name of the user
		Name string
		// Email is the email of the user
		Email string
		// Avatar is the avatar of the user
		Avatar string
		// Bio is the bio of the user
		Bio string
		// Password is the password of the user
		Password string
		// EmailVerifiedAt is timestamp when user email verified
		EmailVerifiedAt *time.Time
		// LockedUntil is temporary lock for brute-force protection
		LockedUntil *time.Time
		// IsActive controls account deactivation
		IsActive bool
	}
	FilterUser struct {
		ID      int
		UUID    string
		Name    string
		Email   string
		Offset  int
		PerPage int
		Sort    string
	}
)

// TableName returns the table name for the user model
//
//	"users"
func (u *User) TableName() string {
	return "users"
}
