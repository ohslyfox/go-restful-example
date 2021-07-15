package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseInstance() *gorm.DB {
	//dsn := "host=127.0.0.1 user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	
	// sqlDB, err := sql.Open("postgres", dsn)
	// db, err := gorm.Open(postgres.New(postgres.Config{
	//   Conn: sqlDB,
	// }), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database connection.")
	}
	db.AutoMigrate(&Book{})
	return db
}
