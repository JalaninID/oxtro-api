package repo_organization

import (
	"app/model"
	"app/pkg"
	"context"
)

// ListOrganization  
func (r *repository) ListOrganization(ctx context.Context, filter model.FilterOrganization) ([]model.Organization, int, error) {
	where := "deleted_at IS NULL "
	args := []any{}

	if filter.Name != "" {
		where, args = pkg.SQLXAndMatch(where, args, "name", filter.Name)
	}
	if filter.UUID != "" {
		where, args = pkg.SQLXAndMatch(where, args, "uuid", filter.UUID)
	}
	if filter.Sort == "" {
		filter.Sort = "created_at DESC"
	}
	if filter.ID != 0 {
		where, args = pkg.SQLXAndMatch(where, args, "id", filter.ID)
	}
	if filter.Domain != "" {
		where, args = pkg.SQLXAndMatch(where, args, "domain", filter.Domain)
	}
	var organizations []model.Organization
	if err := r.db.WithContext(ctx).Scopes(pkg.Paginate(filter.Offset, filter.Limit, filter.Sort, r.db)).Where(where, args...).Find(&organizations).Error; err != nil {
		return nil, 0, err
	}
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Organization{}).Where(where, args...).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return organizations, int(count), nil
}
