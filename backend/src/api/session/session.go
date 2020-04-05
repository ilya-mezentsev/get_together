package session

import (
	"api"
	"github.com/gorilla/mux"
	"interfaces"
	"models"
	"net/http"
)

type Handler struct {
	authService    interfaces.AuthenticationService
	sessionService interfaces.SessionService
}

func InitRequestHandlers(
	authService interfaces.AuthenticationService,
	sessionService interfaces.SessionService,
	middlewares ...mux.MiddlewareFunc,
) {
	handler := Handler{authService, sessionService}
	publicSessionAPI := api.GetRouter().PathPrefix("/session").Subrouter()
	privateSessionAPI := api.GetRouter().PathPrefix("/session").Subrouter()
	for _, middleware := range middlewares {
		privateSessionAPI.Use(middleware)
	}

	publicSessionAPI.HandleFunc("/", handler.getSession).Methods(http.MethodGet)
	publicSessionAPI.HandleFunc("/register", handler.registerUser).Methods(http.MethodPost)
	publicSessionAPI.HandleFunc("/login", handler.loginUser).Methods(http.MethodPost)
	privateSessionAPI.HandleFunc("/user/password", handler.changeUserPassword).Methods(http.MethodPatch)
	privateSessionAPI.HandleFunc("/logout", handler.logoutUser).Methods(http.MethodPost)
}

func (h Handler) getSession(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	s, err := h.sessionService.GetSession(r)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.EncodeAndSendResponse(w, s)
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var registration models.UserCredentials
	api.DecodeRequestBody(r, &registration)

	err := h.authService.RegisterUser(registration)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}

func (h Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var credentials models.UserCredentials
	api.DecodeRequestBody(r, &credentials)

	userSession, err := h.authService.Login(credentials)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	err = h.sessionService.SetSession(r, userSession)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}

func (h Handler) changeUserPassword(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var changePasswordRequest models.ChangePasswordRequest
	api.DecodeRequestBody(r, &changePasswordRequest)

	err := h.authService.ChangePassword(changePasswordRequest.UserId, changePasswordRequest.Password)
	if err != nil {
		panic(api.ApplicationError{OriginalError: err})
	}

	api.SendDefaultResponse(w)
}

func (h Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)
	h.sessionService.InvalidateSession(r)
	api.SendDefaultResponse(w)
}
