package repo_user

import (
	"app/model"
	"context"
)

func (r *repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return user, r.db.WithContext(ctx).Create(&user).Error
}
