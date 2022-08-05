package config

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Migrations struct {
	migrations []*gormigrate.Migration
}

func (m *Migrations) AddMigration(migration *gormigrate.Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrations) AddMigrations(migrations []*gormigrate.Migration) {
	m.migrations = migrations
}

func (m *Migrations) Migrate(db *gorm.DB) error {
	instance := gormigrate.New(db, gormigrate.DefaultOptions, m.migrations)
	if err := instance.Migrate(); err != nil {
		return err
	}
	return nil
}
