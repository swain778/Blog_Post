package database

import (
	"errors"
	"fmt"
	"v0/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=blog port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_ = errors.New("can't connect to database")
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(
		&models.Post{},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}
