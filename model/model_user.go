package model

import (
	"encoding/json"
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
		// UserOrganization is the user organization
		UserOrganization UserOrganization `gorm:"foreignKey:UserID;references:ID"`
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

// UserOrganization represents a user organization
type UserOrganization struct {
	// ID is the primary key
	ID int
	// CreatedAt is the time the user organization was created
	CreatedAt time.Time
	// UpdatedAt is the time the user organization was last updated
	UpdatedAt time.Time
	// DeletedAt is the time the user organization was deleted
	DeletedAt *time.Time
	// UserID is the user ID
	UserID int
	// OrganizationID is the organization ID
	OrganizationID int
	// RoleID is the role ID
	RoleID int
	// Role is the role of the user organization
	Role Role `gorm:"foreignKey:ID;references:RoleID"`
}

// TableName returns the table name for the user organization model
//
//	"user_organizations"
func (u *UserOrganization) TableName() string {
	return "user_organizations"
}

// Role represents a user role
type Role struct {
	//ID is the primary key
	ID int
	//CreatedAt is the time the role was created
	CreatedAt time.Time
	//UpdatedAt is the time the role was last updated
	UpdatedAt time.Time
	//DeletedAt is the time the role was deleted
	DeletedAt *time.Time
	//UUID is the unique identifier for the role
	UUID string
	//OrganizationID is the ID of the organization the role belongs to
	OrganizationID int
	//Name is the name of the role
	Name string
	// Access is the access level of the role
	//
	// Format example:
	//
	// 	[
	// 		{
	// 			"resource": "user",
	// 			"actions": [
	// 				"create",
	// 				"read",
	// 				"update",
	// 				"delete"
	// 			]
	// 		}
	// 	]
	Access string
}

// TableName returns the table name for the role model
//
//	"roles"
func (t *Role) TableName() string {
	return "roles"
}

// Access represents the access level of a role
type Access struct {
	// Resource is the resource the role has access to
	Resource string
	// Actions is the list of actions the role has access to
	Actions []string
}

// GetAccess returns the access level of the role
func (t *Role) GetAccess() []Access {
	var access []Access
	err := json.Unmarshal([]byte(t.Access), &access)
	if err != nil {
		return nil
	}
	return access
}
