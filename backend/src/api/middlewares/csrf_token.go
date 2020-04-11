package middlewares

import (
	"api"
	"encoding/base64"
	"fmt"
	"net/http"
	"plugins/code"
	"regexp"
	"time"
	"utils"
)

const (
	csrfKey = "X-CSRF-Token"
)

var (
	toProtectMethods  = [...]string{http.MethodPost, http.MethodPatch, http.MethodDelete}
	getSessionPathReg = regexp.MustCompile(`.*/session/?$`)
)

type CsrfToken struct {
	PrivateKey string
}

func (c CsrfToken) Check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer api.SendErrorIfPanicked(w)

		if isGetSessionRequest(r) {
			c.setCSRFToken(w)
		}

		if !needToCheckRequest(r.Method) {
			next.ServeHTTP(w, r)
			return
		}

		cookieToken, err := c.getTokenFromCookie(r)
		if err != nil {
			panic(err)
		}

		headerToken, err := c.getTokenFromHeader(r)
		if err != nil {
			panic(err)
		}

		if headerToken == cookieToken {
			next.ServeHTTP(w, r)
			c.setCSRFToken(w)
		} else {
			panic(InvalidCSRFToken)
		}
	})
}

func isGetSessionRequest(r *http.Request) bool {
	return r.Method == http.MethodGet && getSessionPathReg.MatchString(r.URL.Path)
}

func needToCheckRequest(method string) bool {
	for _, m := range toProtectMethods {
		if m == method {
			return true
		}
	}

	return false
}

func (c CsrfToken) getTokenFromCookie(r *http.Request) (string, error) {
	csrfCookie, err := r.Cookie(csrfKey)
	if err != nil {
		return "", NoCSRFCookie
	}

	tokenData, err := code.NewCoder(c.PrivateKey).Decrypt(csrfCookie.Value)
	if err != nil {
		return "", InvalidCSRFCookie
	}

	token, hasToken := tokenData[csrfKey]
	if !hasToken {
		return "", InvalidCSRFCookie
	}

	return token.(string), nil
}

func (c CsrfToken) getTokenFromHeader(r *http.Request) (string, error) {
	encodedToken := r.Header.Get(csrfKey)
	if encodedToken == "" {
		return "", NoCSRFHeader
	}

	token, err := base64.StdEncoding.DecodeString(r.Header.Get(csrfKey))
	if err != nil {
		return "", InvalidCSRFHeader
	}

	return string(token), nil
}

func (c CsrfToken) setCSRFToken(w http.ResponseWriter) {
	publicKey := fmt.Sprintf("%v", time.Now().Unix())
	encodedPublicKey := base64.StdEncoding.EncodeToString([]byte(publicKey))
	tokenData := map[string]interface{}{
		csrfKey: utils.GetHash(c.PrivateKey + publicKey),
	}
	token, err := code.NewCoder(c.PrivateKey).Encrypt(tokenData)
	if err != nil {
		panic(CSRFInternalError)
	}

	w.Header().Add(csrfKey, encodedPublicKey)
	http.SetCookie(w, &http.Cookie{
		Name:     csrfKey,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600,
	})
}
