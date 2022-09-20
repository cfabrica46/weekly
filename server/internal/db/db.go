package db

import (
	"database/sql"
	"fmt"
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
