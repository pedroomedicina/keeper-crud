package config

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"gorm.io/gorm/logger"
	"keeper-crud/helper"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	port, err := strconv.Atoi(os.Getenv("KEEPER_DB_PORT"))
	if err != nil {
		helper.ErrorPanic(err)
	}

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("KEEPER_DB_HOST"),
		port,
		os.Getenv("KEEPER_DB_USER"),
		os.Getenv("KEEPER_DB_PASSWORD"),
		os.Getenv("KEEPER_DB_NAME"),
		os.Getenv("KEEPER_DB_SSLMODE"))
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		helper.ErrorPanic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		helper.ErrorPanic(err)
	}

	gormDB, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	helper.ErrorPanic(err)

	return gormDB
}
