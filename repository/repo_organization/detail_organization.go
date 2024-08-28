package repo_organization

import (
	"app/model"
	"app/pkg"
	"context"
)

func (r *repository) DetailOrganization(ctx context.Context, filter model.FilterOrganization) (model.Organization, error) {
	where := "WHERE deleted_at IS NULL "
	args := []interface{}{}
	if filter.ID != "" {
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
