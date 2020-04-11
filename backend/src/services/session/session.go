package session

import (
	"models"
	"net/http"
	"plugins/code"
	"services/errors"
	"time"
)

const cookieSessionKey = "GT-Session-Token"

type Service struct {
	coder code.Coder
}

func New(key string) Service {
	return Service{code.NewCoder(key)}
}

func (s Service) GetSession(r *http.Request) (map[string]interface{}, error) {
	cookie, err := r.Cookie(cookieSessionKey)
	if err != nil {
		return nil, errors.NoAuthCookie
	}

	decoded, err := s.coder.Decrypt(cookie.Value)
	if err != nil {
		return nil, errors.InvalidAuthCookie
	}

	return decoded, nil
}

func (s Service) SetSession(r *http.Request, session models.UserSession) error {
	token, err := s.coder.Encrypt(map[string]interface{}{
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

func (s Service) InvalidateSession(r *http.Request) {
	r.AddCookie(&http.Cookie{
		Name:     cookieSessionKey,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
