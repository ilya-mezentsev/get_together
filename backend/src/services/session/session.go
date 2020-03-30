package session

import (
	"models"
	"net/http"
	"services/errors"
	"services/session/plugins/code"
	"time"
)

const cookieSessionKey = "GT-Session-Token"

type Service struct {
	coder code.Coder
}

func New(key string) Service {
	return Service{code.NewCoder(key)}
}

func (c Service) GetSession(r *http.Request) (map[string]interface{}, error) {
	cookie, err := r.Cookie(cookieSessionKey)
	if err != nil {
		return nil, errors.NoAuthCookie
	}

	decoded, err := c.coder.Decrypt(cookie.Value)
	if err != nil {
		return nil, errors.InvalidAuthCookie
	}

	return decoded, nil
}

func (c Service) SetSession(r *http.Request, session models.UserSession) error {
	token, err := c.coder.Encrypt(map[string]interface{}{
		"id": session.Id,
	})
	if err != nil {
		return errors.InternalError
	}

	r.AddCookie(&http.Cookie{
		Name:     cookieSessionKey,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600,
	})
	return nil
}

func (c Service) InvalidateSession(r *http.Request) {
	r.AddCookie(&http.Cookie{
		Name:     cookieSessionKey,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
