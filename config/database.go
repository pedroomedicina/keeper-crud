package config

import (
	"fmt"
	"keeper-crud/helper"
	"os"
	"strconv"

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
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
