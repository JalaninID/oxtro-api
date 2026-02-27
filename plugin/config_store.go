package plugin

import (
	"time"

	"gorm.io/gorm"
)

// PluginConfigEntry is the database model for per-plugin configuration.
type PluginConfigEntry struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PluginID  string    `gorm:"uniqueIndex:idx_plugin_config_key;size:255;not null"`
	Key       string    `gorm:"uniqueIndex:idx_plugin_config_key;size:255;not null"`
	Value     string    `gorm:"type:text;not null;default:''"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (PluginConfigEntry) TableName() string {
	return "plugin_configs"
}

// gormPluginConfigStore implements PluginConfigStore using GORM.
type gormPluginConfigStore struct {
	db       *gorm.DB
	pluginID string
}

// NewGormPluginConfigStore creates a config store scoped to a specific plugin.
func NewGormPluginConfigStore(db *gorm.DB, pluginID string) PluginConfigStore {
	return &gormPluginConfigStore{db: db, pluginID: pluginID}
}

func (s *gormPluginConfigStore) Get(key string) (string, error) {
	var entry PluginConfigEntry
	err := s.db.Where("plugin_id = ? AND key = ?", s.pluginID, key).First(&entry).Error
	if err != nil {
		return "", err
	}
	return entry.Value, nil
}

func (s *gormPluginConfigStore) Set(key string, value string) error {
	entry := PluginConfigEntry{
		PluginID:  s.pluginID,
		Key:       key,
		Value:     value,
		UpdatedAt: time.Now(),
	}
	return s.db.
		Where("plugin_id = ? AND key = ?", s.pluginID, key).
		Assign(entry).
		FirstOrCreate(&entry).Error
}

func (s *gormPluginConfigStore) Delete(key string) error {
	return s.db.Where("plugin_id = ? AND key = ?", s.pluginID, key).Delete(&PluginConfigEntry{}).Error
}

func (s *gormPluginConfigStore) GetAll() (map[string]string, error) {
	var entries []PluginConfigEntry
	err := s.db.Where("plugin_id = ?", s.pluginID).Find(&entries).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(entries))
	for _, e := range entries {
		result[e.Key] = e.Value
	}
	return result, nil
}
