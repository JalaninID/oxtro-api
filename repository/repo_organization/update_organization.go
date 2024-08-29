package repo_organization

import (
	"app/model"
	"app/pkg"
	"context"
)

// UpdateOrganization updates an existing organization
func (r *repository) UpdateOrganization(ctx context.Context, filter model.FilterOrganization, organization model.Organization) (model.Organization, error) {
	where := "deleted_at IS NULL "
	args := []any{}
	if filter.Name != "" {
		where, args = pkg.SQLXAndMatch(where, args, "name", filter.Name)
	}
	if filter.ID != 0 {
		where, args = pkg.SQLXAndMatch(where, args, "id", filter.ID)
	}
	if filter.UUID != "" {
		where, args = pkg.SQLXAndMatch(where, args, "uuid", filter.UUID)
	}
	return organization, r.db.WithContext(ctx).Where(where, args...).Updates(&organization).Error
}
