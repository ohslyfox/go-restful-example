package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database connection.")
	}
	db.AutoMigrate(&Book{})
	return db
}
