package config

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Versions struct {
	migrations []*gormigrate.Migration
}

func (v *Versions) AddMigrations(migrations []*gormigrate.Migration) {
	v.migrations = migrations
}

func (v *Versions) Migrate(db *gorm.DB) error {
	instance := gormigrate.New(db, gormigrate.DefaultOptions, v.migrations)
	if err := instance.Migrate(); err != nil {
		return err
	}
	return nil
}
