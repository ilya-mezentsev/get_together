package main

import (
	"api"
	"api/meetings"
	"api/session"
	"api/users"
	"fmt"
	"log"
	"net/http"
	"os"
	"plugins/config"
	"plugins/logger"
	"repositories"
	"services"
	"time"
)

func init() {
	coderKey, err := config.GetCoderKey()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	db, err := config.GetConfiguredConnection()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	meetingsRepository := repositories.Meetings(db)
	meetings.InitRequestHandlers(
		services.Meetings(meetingsRepository),
		services.Participation(repositories.UserSettings(db), repositories.MeetingsSettings(db)),
		services.MeetingsAccessor(meetingsRepository),
	)
	session.InitRequestHandlers(
		services.Authentication(repositories.Credentials(db)),
		services.Session(coderKey),
	)
	users.InitRequestHandlers(services.UserSettings(repositories.UserSettings(db)))
}

func main() {
	logger.Info("Starting application")

	log.Fatal((&http.Server{
		Handler:      api.GetRouter(),
		Addr:         fmt.Sprintf("0.0.0.0:%s", config.GetAPIPort()),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe())
}
