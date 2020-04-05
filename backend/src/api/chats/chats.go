package chats

import (
	"api"
	"github.com/gorilla/mux"
	"interfaces"
	"models"
	"net/http"
	"strconv"
)

type Handler struct {
	chat         interfaces.Chat
	chatAccessor interfaces.ChatAccessor
}

func InitRequestHandlers(
	chat interfaces.Chat,
	chatAccessor interfaces.ChatAccessor,
	middlewares ...mux.MiddlewareFunc,
) {
	handler := Handler{chat, chatAccessor}
	chatsAPI := api.GetRouter().PathPrefix("/chat").Subrouter()
	for _, middleware := range middlewares {
		chatsAPI.Use(middleware)
	}

	chatsAPI.HandleFunc("/meeting/{id:[0-9]+}", handler.getMeetingChat).Methods(http.MethodGet)
	chatsAPI.HandleFunc("/user/{id:[0-9]+}", handler.getUserChats).Methods(http.MethodGet)
	chatsAPI.HandleFunc("/meeting", handler.createMeetingChat).Methods(http.MethodPost)
	chatsAPI.HandleFunc("/meeting/request", handler.createMeetingRequestChat).Methods(http.MethodPost)
	chatsAPI.HandleFunc("/meeting", handler.closeChat).Methods(http.MethodDelete)
	chatsAPI.HandleFunc("/meeting/request", handler.closeChat).Methods(http.MethodDelete)
}

func (h Handler) getMeetingChat(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	// checking of this parameter will be performed in validation proxy
	meetingId, _ := strconv.Atoi(vars["id"])
	chat, err := h.chatAccessor.GetMeetingChat(uint(meetingId))
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, chat)
}

func (h Handler) getUserChats(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	// checking of this parameter will be performed in validation proxy
	userId, _ := strconv.Atoi(vars["id"])
	chats, err := h.chatAccessor.GetUserChats(uint(userId))
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, chats)
}

func (h Handler) createMeetingChat(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.GeneralMeetingRequest
	api.DecodeRequestBody(r, &request)

	err := h.chat.CreateMeetingChat(request.MeetingId)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}

func (h Handler) createMeetingRequestChat(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.GeneralMeetingRequest
	api.DecodeRequestBody(r, &request)

	err := h.chat.CreateMeetingRequestChat(request.MeetingId)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}

func (h Handler) closeChat(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.CloseChatRequest
	api.DecodeRequestBody(r, &request)

	err := h.chat.CloseChat(request.ChatId)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}
