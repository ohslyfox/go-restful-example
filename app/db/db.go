package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bookdb.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open a sqlite db.")
	}
	db.AutoMigrate(&Book{})
	return db
}

func DeleteAllData(db *gorm.DB) {
	db.Exec("DELETE FROM books")
}
