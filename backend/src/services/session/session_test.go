package session

import (
  "fmt"
  "io/ioutil"
  "log"
  mock "mock/services"
  "net/http"
  "os"
  "testing"
  "utils"
)

var sessionController Controller

func init() {
  coderKey := os.Getenv("CODER_KEY")
  if coderKey == "" {
    fmt.Println("CODER_KEY env var is not set")
    os.Exit(1)
  }

  sessionController = New(coderKey)
}

func getRequestWithoutSession() *http.Request {
  req, err := http.NewRequest(http.MethodGet, "", nil)
  if err != nil {
    fmt.Println("should not be error:", err)
    os.Exit(1)
  }

  return req
}

func getRequestWithSession() *http.Request {
  req := getRequestWithoutSession()
  req.AddCookie(&http.Cookie{
    Name: cookieSessionKey,
    Value: mock.TestToken,
    Path: "/",
    HttpOnly: true,
    MaxAge: 3600,
  })
  return req
}

func getRequestWithInvalidSession() *http.Request {
  req := getRequestWithoutSession()
  req.AddCookie(&http.Cookie{
    Name: cookieSessionKey,
    Value: "bad cookie",
  })
  return req
}

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestController_GetSessionSuccess(t *testing.T) {
  session, err := sessionController.GetSession(getRequestWithSession())

  utils.AssertNil(err, t)
  utils.AssertEqual(float64(1), session["id"], t)
}

func TestController_GetSessionNoAuthCookieError(t *testing.T) {
  _, err := sessionController.GetSession(getRequestWithoutSession())

  utils.AssertErrorsEqual(NoAuthCookie, err, t)
}

func TestController_GetSessionInvalidCookieError(t *testing.T) {
  _, err := sessionController.GetSession(getRequestWithInvalidSession())

  utils.AssertErrorsEqual(InvalidAuthCookie, err, t)
}

func TestController_SetSessionSuccess(t *testing.T) {
  req := getRequestWithoutSession()

  err := sessionController.SetSession(req, mock.TestSessionData)

  utils.AssertNil(err, t)

  cookie, err := req.Cookie(cookieSessionKey)
  utils.AssertNil(err, t)
  utils.AssertEqual(mock.TestToken, cookie.Value, t)
}

func TestController_InvalidateSession(t *testing.T) {
  req := getRequestWithSession()

  sessionController.InvalidateSession(req)

  var cookie = &http.Cookie{}
  for _, c := range req.Cookies() {
    if c.Name == cookieSessionKey {
      cookie = c
    }
  }

  utils.AssertEqual("", cookie.Value, t)
}
