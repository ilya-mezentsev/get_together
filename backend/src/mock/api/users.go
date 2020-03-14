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
	UserSettingsResponse struct {
		Status string `json:"status"`
		Data models.FullUserInfo `json:"data"`
	}
)

func FirstUserSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router: r,
		Method: http.MethodGet,
		Endpoint: "user/settings/1",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
	}
}

func NotExistsUserSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router: r,
		Method: http.MethodGet,
		Endpoint: fmt.Sprintf("user/settings/%d", len(repositories.UsersCredentials)+1),
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
	}
}

func InvalidIDUserSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router: r,
		Method: http.MethodGet,
		Endpoint: "user/settings/0",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
	}
}

func PatchFirstUserSettingsRequest(r *mux.Router) utils.RequestData {
	userInfo := repositories.UsersInfo[0]
	return utils.RequestData{
		Router: r,
		Method: http.MethodPatch,
		Endpoint: "user/settings",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
		Data: fmt.Sprintf(
			`{
			"user_id": 1,
			"settings": {"nickname": "hey_sasha_228", "name": "%s", "age": %d, "avatar_url": "http://link.com"}
			}`, userInfo["name"], userInfo["age"]),
	}
}

func PatchNotExistsUserSettingsRequest(r *mux.Router) utils.RequestData {
	userInfo := repositories.UsersInfo[0]
	return utils.RequestData{
		Router: r,
		Method: http.MethodPatch,
		Endpoint: "user/settings",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
		Data: fmt.Sprintf(
			`{
			"user_id": %d,
			"settings": {"nickname": "%s", "name": "%s", "age": %d, "avatar_url": "http://link.com"}
			}`, len(repositories.UsersCredentials)+1, userInfo["nickname"], userInfo["name"], userInfo["age"]),
	}
}

func InvalidIdUserSettingsRequest(r *mux.Router) utils.RequestData {
	userInfo := repositories.UsersInfo[0]
	return utils.RequestData{
		Router: r,
		Method: http.MethodPatch,
		Endpoint: "user/settings",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
		Data: fmt.Sprintf(
			`{
			"user_id": 0, "settings": {"nickname": "%s", "name": "%s", "age": %d, "avatar_url": "http://link.com"}}`,
			userInfo["nickname"], userInfo["name"], userInfo["age"]),
	}
}

func InvalidAllSettingsUserSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router: r,
		Method: http.MethodPatch,
		Endpoint: "user/settings",
		Cookie: &http.Cookie{
			Name: "GT-Session-Token",
			Value: TestToken,
		},
		Data: `{
			"user_id": 1, "settings": {
				"nickname": "hello*world", "name": "bad_<<", "age": 0, "avatar_url": "not_url"
			}}`,
	}
}
