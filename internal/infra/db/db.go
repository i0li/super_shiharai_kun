package db

import (
	"fmt"
	"os"

	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustConnect() *gorm.DB {
	conn, err := newGormPostgres()
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
		panic(err)
	}
	return conn
}

func newGormPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
