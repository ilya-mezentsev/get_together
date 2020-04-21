package messages

import (
	"api"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"interfaces"
	"models"
	"net/http"
	"plugins/logger"
	"strconv"
)

type Handler struct {
	service  interfaces.Messages
	upgrader websocket.Upgrader
}

func InitRequestHandlers(
	service interfaces.Messages,
	middlewares ...mux.MiddlewareFunc,
) {
	handler := Handler{
		service: service,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
	messagesAPI := api.GetRouter().PathPrefix("/messages").Subrouter()
	for _, middleware := range middlewares {
		messagesAPI.Use(middleware)
	}

	api.GetRouter().HandleFunc("/ws", handler.handleWS)
	messagesAPI.HandleFunc(
		"/{chat_id:[0-9]+}/{count:[0-9]+}", handler.getLastMessages).Methods(http.MethodGet)
	messagesAPI.HandleFunc(
		"/{chat_id:[0-9]+}/{message_id:[0-9]+}/{count:[0-9]+}",
		handler.getLastMessagesAfter,
	).Methods(http.MethodGet)
}

func (h Handler) getLastMessages(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	chatId, _ := strconv.Atoi(vars["chat_id"])
	count, _ := strconv.Atoi(vars["count"])

	messages, err := h.service.GetLastMessages(uint(chatId), uint(count))
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, messages)
}

func (h Handler) getLastMessagesAfter(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	chatId, _ := strconv.Atoi(vars["chat_id"])
	messageId, _ := strconv.Atoi(vars["message_id"])
	count, _ := strconv.Atoi(vars["count"])

	messages, err := h.service.GetLastMessagesAfter(uint(chatId), uint(messageId), uint(count))
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, messages)
}

func (h Handler) handleWS(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorF("Error while upgrading connection: %v", err)
	}

	for {
		var message models.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			h.processReadJSONError(conn, err)
			return
		}

		AddConnection(message.ChatId, conn)
		err = h.service.Save(message)
		if err != nil {
			h.processServiceError(conn, err)
			return
		}
		h.sendMessageToAllChatConnections(message)
	}
}

func (h Handler) processReadJSONError(conn *websocket.Conn, err error) {
	logger.ErrorF("Error while reading JSON from connection: %v", err)
	h.trySendErrorMessageToConnection(conn, ReadJSONError)
	RemoveConnection(conn)
	_ = conn.Close()
}

func (h Handler) processServiceError(conn *websocket.Conn, err error) {
	logger.ErrorF("Error while saving message: %v", err)
	h.trySendErrorMessageToConnection(conn, err)
	RemoveConnection(conn)
	_ = conn.Close()
}

func (h Handler) trySendErrorMessageToConnection(conn *websocket.Conn, err error) {
	writeError := conn.WriteJSON(models.ErrorResponse{
		Status:      api.StatusError,
		ErrorDetail: err.Error(),
	})
	if writeError != nil {
		logger.ErrorF("Error while sending message to connection: %v", writeError)
	}
}

func (h Handler) sendMessageToAllChatConnections(message models.Message) {
	for _, connection := range GetConnections(message.ChatId) {
		_ = connection.WriteJSON(message)
	}
}
