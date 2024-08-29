package repo_organization

import (
	"app/model"
	"context"
	"time"
)

// DeleteOrganization deletes an existing organization
func (r *repository) DeleteOrganization(ctx context.Context, filter model.FilterOrganization) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Where("deleted_at IS NULL AND id = ?", filter.ID).Updates(model.Organization{
		DeletedAt: &now,
	}).Error
}
