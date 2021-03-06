package session

import (
	"fmt"
	mock "mock/services"
	"net/http"
	"os"
	"plugins/config"
	"services/errors"
	"testing"
	"utils"
)

var sessionController Service

func init() {
	coderKey, err := config.GetCoderKey()
	if err != nil {
		fmt.Println(err)
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
		Name:     cookieSessionKey,
		Value:    mock.TestToken,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600,
	})
	return req
}

func getRequestWithInvalidSession() *http.Request {
	req := getRequestWithoutSession()
	req.AddCookie(&http.Cookie{
		Name:  cookieSessionKey,
		Value: "bad cookie",
	})
	return req
}

func TestController_GetSessionSuccess(t *testing.T) {
	session, err := sessionController.GetSession(getRequestWithSession())

	utils.AssertNil(err, t)
	utils.AssertEqual(float64(1), session["id"], t)
}

func TestController_GetSessionNoAuthCookieError(t *testing.T) {
	_, err := sessionController.GetSession(getRequestWithoutSession())

	utils.AssertErrorsEqual(errors.NoAuthCookie, err, t)
}

func TestController_GetSessionInvalidCookieError(t *testing.T) {
	_, err := sessionController.GetSession(getRequestWithInvalidSession())

	utils.AssertErrorsEqual(errors.InvalidAuthCookie, err, t)
}

func TestController_SetSessionSuccess(t *testing.T) {
	req := getRequestWithoutSession()

	err := sessionController.SetSession(req, mock.TestSessionData)

	utils.AssertNil(err, t)

	cookie, e := req.Cookie(cookieSessionKey)
	utils.AssertNil(e, t)
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
