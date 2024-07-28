package repo_interest

import (
	"app/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, interest model.Interest) (model.Interest, error) {
	err := r.db.WithContext(ctx).Create(&interest).Error
	if err != nil {
		return model.Interest{}, err
	}
	return interest, err
}
