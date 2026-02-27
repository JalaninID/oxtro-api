package plugin

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// PluginRecord is the database model for persisted plugin state.
type PluginRecord struct {
	ID           int             `gorm:"primaryKey;autoIncrement"`
	PluginID     string          `gorm:"uniqueIndex;size:255;not null"`
	Version      string          `gorm:"size:50;not null"`
	State        PluginState     `gorm:"size:20;not null;default:'installed'"`
	ManifestJSON json.RawMessage `gorm:"column:manifest;type:jsonb;not null;default:'{}'"`
	ConfigJSON   json.RawMessage `gorm:"column:config;type:jsonb;not null;default:'{}'"`
	CreatedAt    time.Time       `gorm:"not null;default:now()"`
	UpdatedAt    time.Time       `gorm:"not null;default:now()"`
}

func (PluginRecord) TableName() string {
	return "plugins"
}

// PluginStore is the repository interface for plugin persistence.
type PluginStore interface {
	SavePluginRecord(ctx context.Context, record PluginRecord) error
	UpdatePluginState(ctx context.Context, pluginID string, state PluginState) error
	DeletePluginRecord(ctx context.Context, pluginID string) error
	ListPluginRecords(ctx context.Context) ([]PluginRecord, error)
	GetPluginRecord(ctx context.Context, pluginID string) (PluginRecord, error)
}

// gormPluginStore implements PluginStore using GORM.
type gormPluginStore struct {
	db *gorm.DB
}

// NewGormPluginStore creates a new GORM-backed plugin store.
func NewGormPluginStore(db *gorm.DB) PluginStore {
	return &gormPluginStore{db: db}
}

func (s *gormPluginStore) SavePluginRecord(ctx context.Context, record PluginRecord) error {
	return s.db.WithContext(ctx).Create(&record).Error
}

func (s *gormPluginStore) UpdatePluginState(ctx context.Context, pluginID string, state PluginState) error {
	return s.db.WithContext(ctx).
		Model(&PluginRecord{}).
		Where("plugin_id = ?", pluginID).
		Updates(map[string]any{
			"state":      state,
			"updated_at": time.Now(),
		}).Error
}

func (s *gormPluginStore) DeletePluginRecord(ctx context.Context, pluginID string) error {
	return s.db.WithContext(ctx).Where("plugin_id = ?", pluginID).Delete(&PluginRecord{}).Error
}

func (s *gormPluginStore) ListPluginRecords(ctx context.Context) ([]PluginRecord, error) {
	var records []PluginRecord
	err := s.db.WithContext(ctx).Find(&records).Error
	return records, err
}

func (s *gormPluginStore) GetPluginRecord(ctx context.Context, pluginID string) (PluginRecord, error) {
	var record PluginRecord
	err := s.db.WithContext(ctx).Where("plugin_id = ?", pluginID).First(&record).Error
	return record, err
}
