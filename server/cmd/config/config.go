package config

import (
	"fmt"
	"os"
)

func SetEnv() (err error) {
	err = os.Setenv("PORT", "8080")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_PORT", "5431")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_USERNAME", "cfabrica46")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_PASSWORD", "01234")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_NAME", "weekly")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	err = os.Setenv("DB_SSLMODE", "disable")
	if err != nil {
		return fmt.Errorf("error to set env: %w", err)
	}

	return nil
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
