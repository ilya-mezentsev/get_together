package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"mock/repositories"
	"models"
	"net/http"
	"utils"
)

type (
	MessagesResponse struct {
		Status string           `json:"status"`
		Data   []models.Message `json:"data"`
	}
)

var (
	DefaultMessagesCount = 2
)

func InvalidProtocolRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Cookie:   cookie,
		Endpoint: "ws",
	}
}

func GetMessagesRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("messages/1/%d", DefaultMessagesCount),
		Cookie:   cookie,
	}
}

func GetMessagesInvalidDataRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "messages/0/0",
		Cookie:   cookie,
	}
}

func GetMessagesAfterMessageRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("messages/1/1/%d", DefaultMessagesCount),
		Cookie:   cookie,
	}
}

func GetMessagesAfterMessageInvalidDataRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "messages/0/0/0",
		Cookie:   cookie,
	}
}

func GetSimpleMessage() models.Message {
	return models.Message{
		ChatId:   1,
		Text:     "Hello",
		SenderId: 1,
	}
}

func GetAnotherSimpleMessage() models.Message {
	return models.Message{
		ChatId:   1,
		Text:     "Hello (2)",
		SenderId: 1,
	}
}

func GetMessageWithNotExistsChatId() models.Message {
	message := GetSimpleMessage()
	message.ChatId = repositories.NotExistsChatId
	return message
}

func GetMessageWithNotExistsSenderId() models.Message {
	message := GetSimpleMessage()
	message.SenderId = repositories.GetNotExistsUserId()
	return message
}

func GetMessageWithInvalidData() models.Message {
	return models.Message{
		ChatId:   0,
		Text:     "",
		SenderId: 0,
	}
}
