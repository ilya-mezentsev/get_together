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

type AllConfigs struct {
	DB             *sqlx.DB
	CoderKey       string
	CsrfPrivateKey string
	Port           string
}

func GetAll() (configs AllConfigs, err error) {
	configs.DB, err = GetConfiguredConnection()
	if err != nil {
		return AllConfigs{}, err
	}

	configs.CoderKey, err = GetCoderKey()
	if err != nil {
		return AllConfigs{}, err
	}

	configs.CsrfPrivateKey, err = GetCSRFPrivateKey()
	if err != nil {
		return AllConfigs{}, err
	}

	configs.Port = GetAPIPort()

	return configs, nil
}

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

func GetCSRFPrivateKey() (string, error) {
	csrfPrivateKey := os.Getenv("CSRF_PRIVATE_KEY")
	if csrfPrivateKey == "" {
		return "", noCSRFPrivateKey
	}

	return csrfPrivateKey, nil
}

func GetAPIPort() string {
	port := os.Getenv("API_PORT")

	if port == "" {
		return port
	} else {
		return fmt.Sprintf("%d", defaultPort)
	}
}
