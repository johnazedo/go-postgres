package gopostgres

import (
	"github.com/johnazedo/gopostgres/config"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	db, err := config.GetInstance()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Migration struct {
	ID       string
	Migrate  func(*gorm.DB) error
	RollBack func(*gorm.DB) error
}
