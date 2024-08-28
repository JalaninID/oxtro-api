package repo_organization

import (
	"app/model"
	"context"
)

// CreateOrganization creates a new organization
func (r *repository) CreateOrganization(ctx context.Context, organization model.Organization) (model.Organization, error) {
	return organization, r.db.WithContext(ctx).Create(&organization).Error
}
