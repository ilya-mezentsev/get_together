package main

import (
	"api"
	"api/session"
	"api/users"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"plugins/logger"
	"repositories"
	"services"
	"time"
)

const (
	defaultPort = 8080
)

func init() {
	coderKey := os.Getenv("CODER_KEY")
	if coderKey == "" {
		fmt.Println("CODER_KEY env var is not set")
		os.Exit(1)
	}

	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		fmt.Println("CONN_STR env var is not set")
		os.Exit(1)
	}

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session.InitRequestHandlers(
		services.Authentication(repositories.Credentials(db)),
		services.Session(coderKey),
	)
	users.InitRequestHandlers(services.UserSettings(repositories.UserSettings(db)))
}

func getApiPort() string {
	port := os.Getenv("API_PORT")

	if port == "" {
		return port
	} else {
		return fmt.Sprintf("%d", defaultPort)
	}
}

func main() {
	logger.Info("Starting application")

	log.Fatal((&http.Server{
		Handler:      api.GetRouter(),
		Addr:         fmt.Sprintf("0.0.0.0:%s", getApiPort()),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe())
}
