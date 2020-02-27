package services

import (
  "fmt"
  "github.com/gorilla/mux"
  "mock/repositories"
  "models"
  "net/http"
  "utils"
)

type (
  DefaultSuccess struct {
    Status string `json:"status"`
    Data interface{} `json:"data"`
  }

  GetSessionSuccess struct {
    Status string `json:"status"`
    Data map[string]interface{} `json:"data"`
  }
)

var (
  TestToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.AfKQ29J6C4MJGnYa0Pw8AnkLdeaAP8ck2BdbyAcoyAc"
  TestSessionData = models.UserSession{ID: 1}
)

func RequestWithSession(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodGet,
    Endpoint: "session/",
    Cookie: &http.Cookie{
      Name: "GT-Session-Token",
      Value: TestToken,
    },
  }
}

func RequestWithoutSession(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodGet,
    Endpoint: "session/",
    Cookie: &http.Cookie{},
  }
}

func RequestWithInvalidToken(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodGet,
    Endpoint: "session/",
    Cookie: &http.Cookie{
      Name: "GT-Session-Token",
      Value: "bad",
    },
  }
}

func SuccessRegistrationRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPost,
    Endpoint: "session/register",
    Data: `{"email": "mather.fucker@gmail.com", "password": "228.me"}`,
    Cookie: &http.Cookie{},
  }
}

func EmailExistsRegistrationRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPost,
    Endpoint: "session/register",
    Data: fmt.Sprintf(`{"email": "%s", "password": "228.me"}`, repositories.UsersCredentials[0]["email"]),
    Cookie: &http.Cookie{},
  }
}

func SuccessLoginRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPost,
    Endpoint: "session/login",
    Data: fmt.Sprintf(`{"email": "%s", "password": "hello"}`, repositories.UsersCredentials[0]["email"]),
    Cookie: &http.Cookie{},
  }
}

func NoCredentialsLoginRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPost,
    Endpoint: "session/login",
    Data: fmt.Sprintf(`{"email": "%s", "password": "not_exists"}`, repositories.UsersCredentials[0]["email"]),
    Cookie: &http.Cookie{},
  }
}

func SuccessChangePasswordRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPatch,
    Endpoint: "session/user/password",
    Data: `{"user_id": 1, "password": "new_password"}`,
    Cookie: &http.Cookie{},
  }
}

func UserIdNotFoundChangePasswordRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPatch,
    Endpoint: "session/user/password",
    Data: fmt.Sprintf(`{"user_id": %d, "password": "new_password"}`, repositories.GetNextUserId()),
    Cookie: &http.Cookie{},
  }
}

func SuccessLogoutRequest(r *mux.Router) utils.RequestData {
  return utils.RequestData{
    Router: r,
    Method: http.MethodPost,
    Endpoint: "session/logout",
    Cookie: &http.Cookie{},
  }
}
