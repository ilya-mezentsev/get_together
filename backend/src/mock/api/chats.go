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
	MeetingChatResponse struct {
		Status string      `json:"status"`
		Data   models.Chat `json:"data"`
	}

	UserChatsResponse struct {
		Status string        `json:"status"`
		Data   []models.Chat `json:"data"`
	}
)

func GetMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/meeting/1",
		Cookie:   &http.Cookie{},
	}
}

func GetMeetingChatByMeetingIfWithoutChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("chat/meeting/%d", repositories.MeetingIdWithoutMeetingChat),
		Cookie:   &http.Cookie{},
	}
}

func GetMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/meeting/0",
		Cookie:   &http.Cookie{},
	}
}

func GetUserChatsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/user/1",
		Cookie:   &http.Cookie{},
	}
}

func GetUserChatsInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/user/0",
		Cookie:   &http.Cookie{},
	}
}

func CreateMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting",
		Cookie:   &http.Cookie{},
		Data:     fmt.Sprintf(`{"meeting_id": %d}`, repositories.MeetingIdWithoutMeetingChat),
	}
}

func CreateMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting",
		Cookie:   &http.Cookie{},
		Data:     `{"meeting_id": 0}`,
	}
}

func CreateMeetingRequestChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting/request",
		Cookie:   &http.Cookie{},
		Data:     `{"meeting_id": 1}`,
	}
}

func CreateMeetingRequestChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting/request",
		Cookie:   &http.Cookie{},
		Data:     `{"meeting_id": 0}`,
	}
}

func CloseMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   &http.Cookie{},
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingChatIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   &http.Cookie{},
		Data:     fmt.Sprintf(`{"chat_id": %d}`, repositories.NotExistsChatId),
	}
}

func CloseMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   &http.Cookie{},
		Data:     `{"chat_id": 0}`,
	}
}

func CloseMeetingRequestChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   &http.Cookie{},
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingRequestChatIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   &http.Cookie{},
		Data:     fmt.Sprintf(`{"chat_id": %d}`, repositories.NotExistsChatId),
	}
}

func CloseMeetingRequestChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   &http.Cookie{},
		Data:     `{"chat_id": 0}`,
	}
}
