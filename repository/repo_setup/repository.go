package repo_setup

import (
	"app/model"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DetailAppSetting(ctx context.Context, key string) (model.AppSetting, error) {
	var setting model.AppSetting
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error
	if err != nil {
		return model.AppSetting{}, err
	}
	return setting, nil
}

func (r *repository) CreateAppSetting(ctx context.Context, setting model.AppSetting) (model.AppSetting, error) {
	return setting, r.db.WithContext(ctx).Create(&setting).Error
}

func (r *repository) CountUsers(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return user, r.db.WithContext(ctx).Create(&user).Error
}
