package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	dsn := os.Getenv("CONNECTION_STRING") //datasource name
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func LoadEnvVariables(paths ...string) {
	var err error
	if len(paths) > 0 {
		err = godotenv.Load(paths...)
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println(".env file successfully loaded")
	}
}

func GetDB() *gorm.DB {
	return db
}
