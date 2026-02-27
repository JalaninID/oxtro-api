package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

// PluginMigrationRecord tracks which migration version a plugin has applied.
type PluginMigrationRecord struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PluginID  string    `gorm:"uniqueIndex;size:255;not null"`
	Version   int       `gorm:"not null;default:0"`
	Dirty     bool      `gorm:"not null;default:false"`
	AppliedAt time.Time `gorm:"not null;default:now()"`
}

func (PluginMigrationRecord) TableName() string {
	return "plugin_migrations"
}

// RunPluginMigrations applies pending up-migrations for a plugin.
func RunPluginMigrations(db *gorm.DB, mp MigrationProvider) error {
	dir := mp.MigrationDir()

	files, err := findMigrationFiles(dir, ".up.sql")
	if err != nil {
		return fmt.Errorf("reading migration dir %s: %w", dir, err)
	}

	if len(files) == 0 {
		return nil
	}

	// Get current version
	var record PluginMigrationRecord
	result := db.Where("plugin_id = ?", mp.MigrationTablePrefix()).First(&record)
	currentVersion := 0
	if result.Error == nil {
		currentVersion = record.Version
	}

	for _, mf := range files {
		if mf.version <= currentVersion {
			continue
		}

		content, err := os.ReadFile(mf.path)
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", mf.path, err)
		}

		if err := db.Exec(string(content)).Error; err != nil {
			// Mark as dirty
			db.Where("plugin_id = ?", mp.MigrationTablePrefix()).
				Assign(PluginMigrationRecord{
					PluginID: mp.MigrationTablePrefix(),
					Version:  mf.version,
					Dirty:    true,
				}).FirstOrCreate(&PluginMigrationRecord{})
			return fmt.Errorf("executing migration %s: %w", mf.path, err)
		}

		// Update version
		db.Where("plugin_id = ?", mp.MigrationTablePrefix()).
			Assign(PluginMigrationRecord{
				PluginID:  mp.MigrationTablePrefix(),
				Version:   mf.version,
				Dirty:     false,
				AppliedAt: time.Now(),
			}).FirstOrCreate(&PluginMigrationRecord{})
	}

	return nil
}

// RollbackPluginMigrations applies all down-migrations for a plugin in reverse order.
func RollbackPluginMigrations(db *gorm.DB, mp MigrationProvider) error {
	dir := mp.MigrationDir()

	files, err := findMigrationFiles(dir, ".down.sql")
	if err != nil {
		return fmt.Errorf("reading migration dir %s: %w", dir, err)
	}

	// Apply in reverse order
	sort.Slice(files, func(i, j int) bool {
		return files[i].version > files[j].version
	})

	for _, mf := range files {
		content, err := os.ReadFile(mf.path)
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", mf.path, err)
		}

		if err := db.Exec(string(content)).Error; err != nil {
			return fmt.Errorf("executing rollback %s: %w", mf.path, err)
		}
	}

	// Remove migration record
	db.Where("plugin_id = ?", mp.MigrationTablePrefix()).Delete(&PluginMigrationRecord{})

	return nil
}

type migrationFile struct {
	version int
	path    string
}

func findMigrationFiles(dir string, suffix string) ([]migrationFile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var files []migrationFile
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), suffix) {
			continue
		}

		var version int
		// Expected format: 000001_description.up.sql
		_, err := fmt.Sscanf(entry.Name(), "%d_", &version)
		if err != nil {
			continue
		}

		files = append(files, migrationFile{
			version: version,
			path:    filepath.Join(dir, entry.Name()),
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].version < files[j].version
	})

	return files, nil
}
