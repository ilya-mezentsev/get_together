package main

import (
	"api"
	"api/chats"
	"api/meetings"
	"api/middlewares"
	"api/session"
	"api/users"
	"fmt"
	"github.com/gorilla/mux"
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
	chatsRepository := repositories.Chat(db)
	sessionService := services.Session(coderKey)
	middlewareFuncs := []mux.MiddlewareFunc{
		middlewares.AuthSession{Service: sessionService}.HasValidSession,
	}

	chats.InitRequestHandlers(
		services.Chat(chatsRepository),
		services.ChatAccessor(chatsRepository),
		middlewareFuncs...,
	)
	meetings.InitRequestHandlers(
		services.Meetings(meetingsRepository),
		services.Participation(repositories.UserSettings(db), repositories.MeetingsSettings(db)),
		services.MeetingsAccessor(meetingsRepository),
		middlewareFuncs...,
	)
	session.InitRequestHandlers(
		services.Authentication(repositories.Credentials(db)),
		sessionService,
		middlewareFuncs...,
	)
	users.InitRequestHandlers(
		services.UserSettings(repositories.UserSettings(db)),
		middlewareFuncs...,
	)
}

func main() {
	addr := fmt.Sprintf("0.0.0.0:%s", config.GetAPIPort())
	logger.Info("Starting application on address " + addr)

	log.Fatal((&http.Server{
		Handler:      api.GetRouter(),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe())
}
