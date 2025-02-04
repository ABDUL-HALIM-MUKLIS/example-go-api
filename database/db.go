package database

import (
	"fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dsn := fmt.Sprintf(
		"host=postgres-container user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}