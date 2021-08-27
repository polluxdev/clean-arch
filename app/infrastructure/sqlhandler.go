package infrastructure

import (
	"fmt"
	"os"

	"github.com/polluxdev/clean-arch/app/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSQLHandler() (interfaces.SQLHandler, error) {
	sqlHandler := interfaces.SQLHandler{}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("HOST_URL"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"), os.Getenv("DB_TIME_ZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlHandler.Conn = db

	return sqlHandler, err
}
