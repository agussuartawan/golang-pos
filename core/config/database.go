package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDatabase() {
	if err := godotenv.Load(".env"); err != nil {
	  log.Fatalf("Error loading .env file")
	}
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	Timezone := os.Getenv("TIMEZONE")

	// connect to database
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=%s connect_timeout=5", DbHost, DbUser, DbPassword, DbName, DbPort, Timezone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Connected to database: ", DbName)
	}
}