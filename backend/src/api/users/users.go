package users

import (
	"api"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"services/user_settings"
	"strconv"
)

type Handler struct {
	usersService user_settings.Service
}

func InitRequestHandlers(usersService user_settings.Service) {
	handler := Handler{usersService}

	usersAPI := api.GetRouter().PathPrefix("/user").Subrouter()

	usersAPI.HandleFunc("/settings/{id:[0-9]+}", handler.getUserSettings).Methods(http.MethodGet)
	usersAPI.HandleFunc("/settings", handler.updateUserSettings).Methods(http.MethodPatch)
}

func (h Handler) getUserSettings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)

	// TODO create service for checking input parameters
	userId, _ := strconv.Atoi(vars["id"])
	info, err := h.usersService.GetUserSettings(uint(userId))
	if err != nil {
		panic(err)
	}

	api.EncodeAndSendResponse(w, info)
}

func (h Handler) updateUserSettings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var updateSettingsRequest models.UpdateUserSettingsRequest
	api.DecodeRequestBody(r, &updateSettingsRequest)

	err := h.usersService.UpdateUserSettings(updateSettingsRequest.UserId, updateSettingsRequest.Settings)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}
