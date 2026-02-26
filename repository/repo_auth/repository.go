package repo_auth

import (
	"app/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateSession(ctx context.Context, session model.AuthSession) (model.AuthSession, error) {
	return session, r.db.WithContext(ctx).Create(&session).Error
}

func (r *repository) DetailSessionByTokenHash(ctx context.Context, tokenHash string) (model.AuthSession, error) {
	var session model.AuthSession
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND deleted_at IS NULL", tokenHash).
		First(&session).Error
	if err != nil {
		return model.AuthSession{}, err
	}
	return session, nil
}

func (r *repository) DetailSessionByID(ctx context.Context, sessionID string, userID int) (model.AuthSession, error) {
	var session model.AuthSession
	err := r.db.WithContext(ctx).
		Where("uuid = ? AND user_id = ? AND deleted_at IS NULL", sessionID, userID).
		First(&session).Error
	if err != nil {
		return model.AuthSession{}, err
	}
	return session, nil
}

func (r *repository) ListSessions(ctx context.Context, userID int) ([]model.AuthSession, error) {
	var sessions []model.AuthSession
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *repository) RevokeSessionByTokenHash(ctx context.Context, tokenHash string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.AuthSession{}).
		Where("token_hash = ? AND deleted_at IS NULL AND revoked_at IS NULL", tokenHash).
		Update("revoked_at", now).Error
}

func (r *repository) RevokeSessionByID(ctx context.Context, sessionID string, userID int) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.AuthSession{}).
		Where("uuid = ? AND user_id = ? AND deleted_at IS NULL AND revoked_at IS NULL", sessionID, userID).
		Update("revoked_at", now).Error
}

func (r *repository) RevokeAllSessionsByUserID(ctx context.Context, userID int) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.AuthSession{}).
		Where("user_id = ? AND deleted_at IS NULL AND revoked_at IS NULL", userID).
		Update("revoked_at", now).Error
}

func (r *repository) RotateSession(ctx context.Context, oldSessionID int, session model.AuthSession) (model.AuthSession, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return model.AuthSession{}, tx.Error
	}

	now := time.Now()
	if err := tx.Model(&model.AuthSession{}).
		Where("id = ? AND revoked_at IS NULL", oldSessionID).
		Update("revoked_at", now).Error; err != nil {
		tx.Rollback()
		return model.AuthSession{}, err
	}

	if err := tx.Create(&session).Error; err != nil {
		tx.Rollback()
		return model.AuthSession{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.AuthSession{}, err
	}

	return session, nil
}

func (r *repository) CreatePasswordReset(ctx context.Context, reset model.PasswordReset) (model.PasswordReset, error) {
	return reset, r.db.WithContext(ctx).Create(&reset).Error
}

func (r *repository) DetailPasswordReset(ctx context.Context, tokenHash string) (model.PasswordReset, error) {
	var reset model.PasswordReset
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND deleted_at IS NULL", tokenHash).
		First(&reset).Error
	if err != nil {
		return model.PasswordReset{}, err
	}
	return reset, nil
}

func (r *repository) MarkPasswordResetUsed(ctx context.Context, resetID int) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.PasswordReset{}).
		Where("id = ? AND used_at IS NULL", resetID).
		Update("used_at", now).Error
}

func (r *repository) UpsertEmailVerification(ctx context.Context, verification model.EmailVerification) (model.EmailVerification, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return model.EmailVerification{}, tx.Error
	}

	if err := tx.Where("user_id = ? AND used_at IS NULL", verification.UserID).
		Delete(&model.EmailVerification{}).Error; err != nil {
		tx.Rollback()
		return model.EmailVerification{}, err
	}

	if err := tx.Create(&verification).Error; err != nil {
		tx.Rollback()
		return model.EmailVerification{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.EmailVerification{}, err
	}
	return verification, nil
}

func (r *repository) DetailEmailVerification(ctx context.Context, tokenHash string) (model.EmailVerification, error) {
	var verification model.EmailVerification
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND deleted_at IS NULL", tokenHash).
		First(&verification).Error
	if err != nil {
		return model.EmailVerification{}, err
	}
	return verification, nil
}

func (r *repository) MarkEmailVerificationUsed(ctx context.Context, verificationID int) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.EmailVerification{}).
		Where("id = ? AND used_at IS NULL", verificationID).
		Update("used_at", now).Error
}

func (r *repository) CreateLoginAttempt(ctx context.Context, attempt model.LoginAttempt) (model.LoginAttempt, error) {
	return attempt, r.db.WithContext(ctx).Create(&attempt).Error
}

func (r *repository) CountRecentFailedLoginAttempts(ctx context.Context, userID int, sinceMinutes int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.LoginAttempt{}).
		Where("user_id = ? AND success = false AND created_at >= ? AND deleted_at IS NULL", userID, time.Now().Add(-time.Duration(sinceMinutes)*time.Minute)).
		Count(&count).Error
	return count, err
}

func (r *repository) CreateSecurityEvent(ctx context.Context, event model.SecurityEvent) (model.SecurityEvent, error) {
	return event, r.db.WithContext(ctx).Create(&event).Error
}
