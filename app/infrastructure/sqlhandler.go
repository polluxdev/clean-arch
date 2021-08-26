package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSQLHandler() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("HOST_URL"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"), os.Getenv("DB_TIME_ZONE"),
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connected to database")
	}

	log.Print("Connected to database")

	return db
}
