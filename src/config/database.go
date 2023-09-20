package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/obskur123/crent/src/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

func CreateConnection() *gorm.DB {

	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Buenos_Aires",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Singleton pattern applied
	once.Do(func() {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	helper.ErrorPanic(err)

	return db

}