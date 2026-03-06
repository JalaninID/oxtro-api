package service_setup

import (
	"app/constant"
	"app/domain"
	"app/dto/dto_setup"
	"app/model"
	"app/pkg"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const installationCompletedKey = "installation_completed_at"

type service struct {
	repoSetup domain.RepositorySetup
	db        *gorm.DB
	logger    *logrus.Logger
}

func NewService(repoSetup domain.RepositorySetup, db *gorm.DB, logger *logrus.Logger) *service {
	return &service{
		repoSetup: repoSetup,
		db:        db,
		logger:    logger,
	}
}

func (s *service) GetStatus(ctx context.Context) (*dto_setup.SetupStatusResponse, error) {
	setting, err := s.repoSetup.DetailAppSetting(ctx, installationCompletedKey)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &dto_setup.SetupStatusResponse{
			IsCompleted: false,
			DBHealthy:   s.checkDatabaseHealth(ctx),
		}, nil
	}
	if err != nil {
		s.logger.Errorf("setup status detail app setting failed: %v", err)
		return nil, err
	}

	return &dto_setup.SetupStatusResponse{
		IsCompleted: true,
		DBHealthy:   s.checkDatabaseHealth(ctx),
		CompletedAt: setting.Value,
	}, nil
}

func (s *service) IsSetupCompleted(ctx context.Context) (bool, error) {
	status, err := s.GetStatus(ctx)
	if err != nil {
		s.logger.Errorf("setup status check failed: %v", err)
		return false, err
	}
	return status.IsCompleted, nil
}

func (s *service) RunSetup(ctx context.Context, req dto_setup.RunSetupRequest) (*dto_setup.RunSetupResponse, error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, constant.ErrInvalidCredentials
	}
	if err := pkg.ValidatePassword(req.Password); err != nil {
		return nil, constant.ErrPasswordPolicy
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := txRepository{db: tx}

		if _, err := txRepo.DetailAppSetting(ctx, installationCompletedKey); err == nil {
			return constant.ErrSetupAlreadyCompleted
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		userCount, err := txRepo.CountUsers(ctx)
		if err != nil {
			return err
		}
		if userCount > 0 {
			return constant.ErrSetupAlreadyCompleted
		}

		passwordHash, err := pkg.HashPassword(req.Password)
		if err != nil {
			return err
		}

		now := time.Now().UTC()
		_, err = txRepo.CreateUser(ctx, model.User{
			UUID:            uuid.NewString(),
			Name:            req.Username,
			Email:           req.Email,
			Password:        passwordHash,
			IsActive:        true,
			EmailVerifiedAt: &now,
		})
		if err != nil {
			s.logger.Errorf("setup create bootstrap user failed: %v", err)
			return err
		}

		setting := model.AppSetting{
			Key:   installationCompletedKey,
			Value: now.Format(time.RFC3339),
		}
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoNothing: true,
		}).Create(&setting)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return constant.ErrSetupAlreadyCompleted
		}

		return nil
	})
	if err != nil {
		s.logger.Errorf("run setup failed: %v", err)
		return nil, err
	}

	return &dto_setup.RunSetupResponse{
		Success: true,
		Message: "setup completed",
	}, nil
}

func (s *service) checkDatabaseHealth(ctx context.Context) bool {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.logger.Warnf("setup db health db instance unavailable: %v", err)
		return false
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		s.logger.Warnf("setup db ping failed: %v", err)
		return false
	}
	return true
}

type txRepository struct {
	db *gorm.DB
}

func (r txRepository) DetailAppSetting(ctx context.Context, key string) (model.AppSetting, error) {
	var setting model.AppSetting
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error
	if err != nil {
		return model.AppSetting{}, err
	}
	return setting, nil
}

func (r txRepository) CountUsers(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r txRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return user, r.db.WithContext(ctx).Create(&user).Error
}
