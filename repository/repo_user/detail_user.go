package repo_user

import (
	"app/model"
	"app/pkg"
	"context"
)

func (r *repository) DetailUser(ctx context.Context, filter model.FilterUser) (model.User, error) {
	where := " deleted_at IS NULL "
	args := []any{}
	if filter.ID != 0 {
		where, args = pkg.SQLXAndMatch(where, args, "id", filter.ID)
	}
	if filter.Email != "" {
		where, args = pkg.SQLXAndMatch(where, args, "email", filter.Email)
	}
	if filter.UUID != "" {
		where, args = pkg.SQLXAndMatch(where, args, "uuid", filter.UUID)
	}

	var user model.User
	err := r.db.WithContext(ctx).Where(where, args...).
		Preload("UserOrganization").Preload("UserOrganization.Role").
		First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
