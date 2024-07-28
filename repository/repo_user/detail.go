package repo_user

import (
	"app/model"
	"context"
)

func (r *Repository) DetailByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// detail by username
func (r *Repository) DetailByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND username = ?", username).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
