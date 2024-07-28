package repo_interest

import (
	"app/model"
	"time"
)

func (r *Repository) DeletePartnerExpired() error {
	r.logger.Info("Start Cron: DeletePartnerExpired")
	var (
		interest model.Interest
		where    string = "deleted_at IS NULL "
	)
	yesterday := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	where += " AND created_at < ?"
	err := r.db.Model(&model.Interest{}).Where(where, yesterday).Delete(&interest).Error
	if err != nil {
		r.logger.Error(err)
		return err
	}
	r.logger.Info("End Cron: DeletePartnerExpired")
	return nil
}
