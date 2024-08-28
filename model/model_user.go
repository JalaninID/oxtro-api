package model

import (
	"time"
)

// User represents a user
type User struct {
	// ID is the primary key
	ID        int
	// CreatedAt is the time the user was created
	CreatedAt time.Time
	// UpdatedAt is the time the user was last updated
	UpdatedAt time.Time
	// DeletedAt is the time the user was deleted
	DeletedAt *time.Time
	// UUID is the unique identifier for the user
	UUID      string
	// Name is the name of the user
	Name      string
	// Email is the email of the user
	Email     string
	// Avatar is the avatar of the user
	Avatar    string
	// Bio is the bio of the user
	Bio       string
	// Password is the password of the user
	Password  string
}

// TableName returns the table name for the user model
//	"users"
func (u *User) TableName() string {
	return "users"
}

// UserOrganization represents a user organization
type UserOrganization struct {
	// ID is the primary key
	ID             int
	// CreatedAt is the time the user organization was created
	CreatedAt      time.Time
	// UpdatedAt is the time the user organization was last updated
	UpdatedAt      time.Time
	// DeletedAt is the time the user organization was deleted
	DeletedAt      *time.Time
	// UserID is the user ID
	UserID         int
	// OrganizationID is the organization ID
	OrganizationID int
	// RoleID is the role ID
	RoleID         int
}

// TableName returns the table name for the user organization model
//	"user_organizations"
func (u *UserOrganization) TableName() string {
	return "user_organizations"
}