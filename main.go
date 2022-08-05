package gopostgres

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/johnazedo/gopostgres/config"
	"gorm.io/gorm"
)

func SetupPostgres(migrations []*Migration) (*Database, error) {
	db, err := config.GetInstance()
	if err != nil {
		return nil, err
	}
	m := config.Migrations{}

	for _, migration := range migrations {
		m.AddMigration(&migration.Migration)
	}

	err = m.Migrate(db)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

type Migration struct {
	gormigrate.Migration
}

type Database struct {
	instance *gorm.DB
}