package repo_user

import (
	"app/model"
	"app/pkg"
)

func (r *Repository) SearchPartner(filter model.FilterInterest) ([]model.User, int, error) {
	var (
		users     []model.User
		args      []any
		where     string = "deleted_at IS NULL "
		totalRows int64
	)

	if filter.Gender != "" {
		where, args = pkg.SQLXAndMatch(where, args, "gender", filter.Gender)
	}
	if len(filter.InterestUserID) > 0 {
		where, args = pkg.SQLXAndNotIn(where, args, "id", filter.InterestUserID)
	}

	err := r.db.Scopes(
		pkg.Paginate(filter.Offset, filter.PerPage, filter.GetSort(), r.db)).
		Where(where, args...).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	if err := r.db.Model([]model.User{}).Where(where, args...).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}
	return users, int(totalRows), nil
}
