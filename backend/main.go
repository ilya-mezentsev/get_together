package main

import (
	"api"
	"api/chats"
	"api/meetings"
	"api/messages"
	"api/middlewares"
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

var (
	addr string
)

func init() {
	configs, err := config.GetAll()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	addr = fmt.Sprintf("0.0.0.0:%s", configs.Port)
	meetingsRepository := repositories.Meetings(configs.DB)
	chatsRepository := repositories.Chat(configs.DB)
	sessionService := services.Session(configs.CoderKey)
	checkSessionMiddleware := middlewares.AuthSession{Service: sessionService}.HasValidSession

	api.GetRouter().Use(middlewares.CsrfToken{PrivateKey: configs.CsrfPrivateKey}.Check)
	chats.InitRequestHandlers(
		services.Chat(chatsRepository),
		services.ChatAccessor(chatsRepository),
		checkSessionMiddleware,
	)
	meetings.InitRequestHandlers(
		services.Meetings(meetingsRepository),
		services.Participation(repositories.UserSettings(configs.DB), repositories.MeetingsSettings(configs.DB)),
		services.MeetingsAccessor(meetingsRepository),
		checkSessionMiddleware,
	)
	messages.InitRequestHandlers(
		services.Messages(repositories.Messages(configs.DB)),
		checkSessionMiddleware,
	)
	session.InitRequestHandlers(
		services.Authentication(repositories.Credentials(configs.DB)),
		sessionService,
		checkSessionMiddleware,
	)
	users.InitRequestHandlers(
		services.UserSettings(repositories.UserSettings(configs.DB)),
		checkSessionMiddleware,
	)
}

func main() {
	logger.Info("Starting application on address " + addr)

	log.Fatal((&http.Server{
		Handler:      api.GetRouter(),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe())
}
