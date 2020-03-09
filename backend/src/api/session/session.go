package session

import (
  "api"
  "models"
  "net/http"
  "plugins/logger"
  "services/authentication"
  "services/session"
)

type Handler struct {
  authService authentication.Service
  sessionService session.Service
}

func InitRequestHandlers(authService authentication.Service, sessionService session.Service) {
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

  s, eWrapper := h.sessionService.GetSession(r)
  if eWrapper != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while getting session: %v; returning: %v",
      Args: []interface{}{eWrapper.OriginalError(), eWrapper.ExternalError()},
    }, logger.Warning)

    panic(eWrapper.ExternalError())
  }

  api.EncodeAndSendResponse(w, s)
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var registration models.UserCredentials
  api.DecodeRequestBody(r, &registration)

  eWrapper := h.authService.RegisterUser(registration)
  if eWrapper != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while register user: %v",
      Args: []interface{}{eWrapper.OriginalError()},
      Optional: map[string]interface{}{
        "credentials": registration,
      },
    }, logger.Warning)

    panic(eWrapper.ExternalError())
  }

  api.SendDefaultResponse(w)
}

func (h Handler) loginUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var credentials models.UserCredentials
  api.DecodeRequestBody(r, &credentials)

  userSession, err := h.authService.Login(credentials)
  if err != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while login user: %v",
      Args: []interface{}{err.OriginalError()},
      Optional: map[string]interface{}{
        "credentials": credentials,
      },
    }, logger.Warning)

    panic(err.ExternalError())
  }

  h.sessionService.SetSession(r, userSession)
  api.SendDefaultResponse(w)
}

func (h Handler) changeUserPassword(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)

  var changePasswordRequest models.ChangePasswordRequest
  api.DecodeRequestBody(r, &changePasswordRequest)

  err := h.authService.ChangePassword(changePasswordRequest.UserId, changePasswordRequest.Password)
  if err != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while changing password: %v",
      Args: []interface{}{err.OriginalError()},
      Optional: map[string]interface{}{
        "request": changePasswordRequest,
      },
    }, logger.Warning)

    panic(err.ExternalError())
  }

  api.SendDefaultResponse(w)
}

func (h Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
  defer api.SendErrorIfPanicked(w)
  h.sessionService.InvalidateSession(r)
  api.SendDefaultResponse(w)
}
