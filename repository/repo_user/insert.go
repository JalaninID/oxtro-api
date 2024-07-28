package repo_user

import (
	"app/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
