package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
	"fmt"
	"time"
)

var DB *gorm.DB
var err error

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	dsn = fmt.Sprintf(dsn, dbUser, dbPassword, dbHost, dbPort, dbName)

	maxAttempts := 5
	delay := 3 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Attempting to connect to the database (Attempt %d/%d)\n", attempts, maxAttempts)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			sqlDB, _ := DB.DB()
			err = sqlDB.Ping()
		}

		if err == nil {
			fmt.Println("Successfully connected to the database.")
			break
		}
		
		fmt.Printf("Failed to connect to the database: %v\n", err)
		if err != nil && attempts < maxAttempts {
			fmt.Printf("Retrying in %v...\n", delay)
			time.Sleep(delay)
		} else {
			panic ("Could not connect to the database after several attempts.")
		}
	}
}