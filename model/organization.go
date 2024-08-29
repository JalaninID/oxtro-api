package model

import "time"

// Organization represents an organization
type (
	Organization struct {
		// ID is the primary key
		ID int
		// CreatedAt is the time the organization was created
		CreatedAt time.Time
		// UpdatedAt is the time the organization was last updated
		UpdatedAt time.Time
		// DeletedAt is the time the organization was deleted
		DeletedAt *time.Time
		// UUID is the unique identifier for the organization
		UUID string
		// Name is the name of the organization
		Name string
		// Domain is the domain of the organization
		Domain string
		// Logo is the logo of the organization
		Logo string
		// Description is the description of the organization
		Description string
	}
	FilterOrganization struct {
		ID     int
		UUID   string
		Name   string
		Domain string
		Offset int
		Limit  int
		Sort   string
	}
)

// TableName returns the table name for the organization model
//
//	"organizations"
func (t *Organization) TableName() string {
	return "organizations"
}
