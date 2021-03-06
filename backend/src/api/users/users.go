package users

import (
	"api"
	"github.com/gorilla/mux"
	"interfaces"
	"models"
	"net/http"
	"strconv"
)

type Handler struct {
	usersService interfaces.UsersSettings
}

func InitRequestHandlers(
	usersService interfaces.UsersSettings,
	middlewares ...mux.MiddlewareFunc,
) {
	handler := Handler{usersService}
	usersAPI := api.GetRouter().PathPrefix("/user").Subrouter()
	for _, middleware := range middlewares {
		usersAPI.Use(middleware)
	}

	usersAPI.HandleFunc("/settings/{id:[0-9]+}", handler.getUserSettings).Methods(http.MethodGet)
	usersAPI.HandleFunc("/settings", handler.updateUserSettings).Methods(http.MethodPatch)
}

func (h Handler) getUserSettings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	// checking of this parameter will be performed in validation proxy
	userId, _ := strconv.Atoi(vars["id"])
	info, err := h.usersService.GetUserSettings(uint(userId))
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, info)
}

func (h Handler) updateUserSettings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var updateSettingsRequest models.UpdateUserSettingsRequest
	api.DecodeRequestBody(r, &updateSettingsRequest)

	err := h.usersService.UpdateUserSettings(updateSettingsRequest.UserId, updateSettingsRequest.Settings)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}
