package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var instance *gorm.DB
var locker = &sync.Mutex{}

func GetInstance() (*gorm.DB, error) {
	info := getDatabaseInfo()
	locker.Lock()
	if instance == nil {
		var err error
		instance, err = gorm.Open(postgres.Open(info), &gorm.Config{})
		if err != nil {
			locker.Unlock()
			return nil, err
		}
	}
	locker.Unlock()
	return instance, nil
}

func getDatabaseInfo() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
}
