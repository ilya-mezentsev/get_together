package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"plugins/logger"
	"time"
)

const (
	defaultPort     = 8080
	maxOpenConns    = 30
	maxIdleConns    = 30
	connMaxLifetime = time.Hour
)

func GetConfiguredConnection() (*sqlx.DB, error) {
	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		return nil, noConnectionString
	}

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.ErrorF("Error while opening DB: %v", err)
		return nil, cannotOpenDB
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}

func GetCoderKey() (string, error) {
	coderKey := os.Getenv("CODER_KEY")
	if coderKey == "" {
		return "", noCoderKey
	}

	return coderKey, nil
}

func GetAPIPort() string {
	port := os.Getenv("API_PORT")

	if port == "" {
		return port
	} else {
		return fmt.Sprintf("%d", defaultPort)
	}
}
