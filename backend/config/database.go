package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"toko-opname/models"

	"os"
	"log"
	"time"
	"fmt"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	dsn = fmt.Sprintf(dsn, dbUser, dbPassword, dbHost, dbPort, dbName)

	maxAttempts := 5
	delay := 3 * time.Second

	var err error
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Attempting connect to database... attempt: %d/%d", attempts, maxAttempts)
			log.Printf("Retrying in %s...", delay)
			time.Sleep(delay)
		} else {
			log.Println("Connected to database successfully")
			break
		}
	}

	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	return DB
}