package repo_interest

import (
	"app/model"
	"app/pkg"
)

func (r *Repository) ListPartnerSwipe(filter model.FilterInterest) ([]string, error) {
	var (
		userIDs []string
		args    []any
		where   string = "deleted_at IS NULL "
	)
	if filter.UserID != "" {
		where, args = pkg.SQLXAndMatch(where, args, "user_id", filter.UserID)
	}

	err := r.db.Model(&model.Interest{}).Where(where, args...).Pluck("interest_user_id", &userIDs).Error
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}
