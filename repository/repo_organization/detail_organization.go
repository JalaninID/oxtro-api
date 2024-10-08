package repo_organization

import (
	"app/model"
	"app/pkg"
	"context"
)

// DetailOrganization
func (r *repository) DetailOrganization(ctx context.Context, filter model.FilterOrganization) (model.Organization, error) {
	where := " deleted_at IS NULL "
	args := []any{}
	if filter.ID != 0 {
		where, args = pkg.SQLXAndMatch(where, args, "id", filter.ID)
	}
	if filter.UUID != "" {
		where, args = pkg.SQLXAndMatch(where, args, "uuid", filter.UUID)
	}
	if filter.Domain != "" {
		where, args = pkg.SQLXAndMatch(where, args, "domain", filter.Domain)
	}
	var organization model.Organization
	if err := r.db.WithContext(ctx).Where(where, args...).First(&organization).Error; err != nil {
		return model.Organization{}, err
	}
	return organization, nil
}
