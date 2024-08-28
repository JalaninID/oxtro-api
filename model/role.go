package model

import (
	"encoding/json"
	"time"
)

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
