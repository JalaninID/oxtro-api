package repo_user

import (
	"app/model"
	"context"
	"time"
)

func (r *Repository) UpdatePurchasePremium(ctx context.Context, userID string, premium bool) error {
	return r.db.Model(&model.User{}).Where("id = ?", userID).Updates(&model.User{
		UpdatedAt: time.Now().UTC(),
		IsPremium: premium,
	}).Error
}
