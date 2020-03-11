package session

import (
  "api"
  "interfaces"
  "models"
  "net/http"
)

type Handler struct {
  authService interfaces.AuthenticationService
  sessionService interfaces.SessionService
}

func InitRequestHandlers(authService interfaces.AuthenticationService, sessionService interfaces.SessionService) {
  handler := Handler{authService, sessionService}

  sessionAPI := api.GetRouter().PathPrefix("/session").Subrouter()

  sessionAPI.HandleFunc("/", handler.getSession).Methods(http.MethodGet)
  sessionAPI.HandleFunc("/register", handler.registerUser).Methods(http.MethodPost)
  sessionAPI.HandleFunc("/login", handler.loginUser).Methods(http.MethodPost)
  sessionAPI.HandleFunc("/user/password", handler.changeUserPassword).Methods(http.MethodPatch)
  sessionAPI.HandleFunc("/logout", handler.logoutUser).Methods(http.MethodPost)
}

func (h Handler) getSession(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  s, err := h.sessionService.GetSession(r)
  if err != nil {
    panic(err)
  }

  api.EncodeAndSendResponse(w, s)
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var registration models.UserCredentials
  api.DecodeRequestBody(r, &registration)

  err := h.authService.RegisterUser(registration)
  if err != nil {
    panic(err)
  }

  api.SendDefaultResponse(w)
}

func (h Handler) loginUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var credentials models.UserCredentials
  api.DecodeRequestBody(r, &credentials)

  userSession, err := h.authService.Login(credentials)
  if err != nil {
    panic(err)
  }

  err = h.sessionService.SetSession(r, userSession)
  if err != nil {
    panic(err)
  }

  api.SendDefaultResponse(w)
}

func (h Handler) changeUserPassword(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var changePasswordRequest models.ChangePasswordRequest
  api.DecodeRequestBody(r, &changePasswordRequest)

  err := h.authService.ChangePassword(changePasswordRequest.UserId, changePasswordRequest.Password)
  if err != nil {
    panic(err)
  }

  api.SendDefaultResponse(w)
}

func (h Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)
  h.sessionService.InvalidateSession(r)
  api.SendDefaultResponse(w)
}
