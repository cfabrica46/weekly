package db

import (
	"database/sql"
	"fmt"
	"os"
	// _ "github.com/lib/pq".
)

type DB struct {
	*sql.DB
}

const DriverDefault = "postgres"

func Open(driver, dbInfo string) (db *DB, err error) {
	initDB, err := sql.Open(driver, dbInfo)
	if err != nil {
		return nil, fmt.Errorf("error to connect db: %w", err)
	}

	err = initDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error to connect db: %w", err)
	}

	return &DB{initDB}, nil
}

func GetDBInfo() (dbInfo string) {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}
