package repo_user

import (
	"app/model"
	"app/pkg"
	"context"
)

func (r *repository) UpdateUser(ctx context.Context, filter model.FilterUser, data model.User) (model.User, error) {
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

	err := r.db.WithContext(ctx).Where(where, args...).Updates(&data).Error
	if err != nil {
		return model.User{}, err
	}
	return data, nil
}
