package session

import (
  "fmt"
  "io/ioutil"
  "log"
  "mock"
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

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(float64(1) == session["id"], func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: 1, Got: session["id"]}))
    t.Fail()
  })
}

func TestController_GetSessionNoAuthCookieError(t *testing.T) {
  _, err := sessionController.GetSession(getRequestWithoutSession())

  utils.AssertErrorsEqual(NoAuthCookie, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestController_GetSessionInvalidCookieError(t *testing.T) {
  _, err := sessionController.GetSession(getRequestWithInvalidSession())

  utils.AssertErrorsEqual(InvalidAuthCookie, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestController_SetSessionSuccess(t *testing.T) {
  req := getRequestWithoutSession()

  err := sessionController.SetSession(req, mock.TestSessionData)

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })

  cookie, err := req.Cookie(cookieSessionKey)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.TestToken == cookie.Value, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.TestToken, Got: cookie.Value}))
    t.Fail()
  })
}

func TestController_InvalidateSession(t *testing.T) {
  req := getRequestWithSession()

  sessionController.InvalidateSession(req)

  var cookie *http.Cookie
  for _, c := range req.Cookies() {
    if c.Name == cookieSessionKey {
      cookie = c
    }
  }

  utils.Assert(nil != cookie && "" == cookie.Value, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "", Got: cookie.Value}))
    t.Fail()
  })
}
