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

var (
	cookie = &http.Cookie{
		Name:  "GT-Session-Token",
		Value: TestToken,
	}
	emptyCookie = &http.Cookie{}
)

func GetMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/meeting/1",
		Cookie:   cookie,
	}
}

func GetMeetingChatRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/meeting/1",
		Cookie:   emptyCookie,
	}
}

func GetMeetingChatByMeetingIfWithoutChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("chat/meeting/%d", repositories.MeetingIdWithoutMeetingChat),
		Cookie:   cookie,
	}
}

func GetMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/meeting/0",
		Cookie:   cookie,
	}
}

func GetUserChatsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/user/1",
		Cookie:   cookie,
	}
}

func GetUserChatsWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/user/1",
		Cookie:   emptyCookie,
	}
}

func GetUserChatsInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "chat/user/0",
		Cookie:   cookie,
	}
}

func CreateMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting",
		Cookie:   cookie,
		Data:     fmt.Sprintf(`{"meeting_id": %d}`, repositories.MeetingIdWithoutMeetingChat),
	}
}

func CreateMeetingChatRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting",
		Cookie:   emptyCookie,
		Data:     fmt.Sprintf(`{"meeting_id": %d}`, repositories.MeetingIdWithoutMeetingChat),
	}
}

func CreateMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting",
		Cookie:   cookie,
		Data:     `{"meeting_id": 0}`,
	}
}

func CreateMeetingRequestChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting/request",
		Cookie:   cookie,
		Data:     `{"meeting_id": 1}`,
	}
}

func CreateMeetingRequestChatRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting/request",
		Cookie:   emptyCookie,
		Data:     `{"meeting_id": 1}`,
	}
}

func CreateMeetingRequestChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "chat/meeting/request",
		Cookie:   cookie,
		Data:     `{"meeting_id": 0}`,
	}
}

func CloseMeetingChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   cookie,
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingChatRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   emptyCookie,
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingChatIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   cookie,
		Data:     fmt.Sprintf(`{"chat_id": %d}`, repositories.NotExistsChatId),
	}
}

func CloseMeetingChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting",
		Cookie:   cookie,
		Data:     `{"chat_id": 0}`,
	}
}

func CloseMeetingRequestChatRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   cookie,
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingRequestChatRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   emptyCookie,
		Data:     `{"chat_id": 1}`,
	}
}

func CloseMeetingRequestChatIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   cookie,
		Data:     fmt.Sprintf(`{"chat_id": %d}`, repositories.NotExistsChatId),
	}
}

func CloseMeetingRequestChatInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "chat/meeting/request",
		Cookie:   cookie,
		Data:     `{"chat_id": 0}`,
	}
}
