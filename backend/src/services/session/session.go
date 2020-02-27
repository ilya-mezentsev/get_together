package session

import (
  "interfaces"
  "models"
  "net/http"
  "services"
  "services/session/plugins/code"
  "time"
)

const cookieSessionKey = "GT-Session-Token"

type Service struct {
  coder code.Coder
}

func New(key string) Service {
  return Service{coder: code.NewCoder(key)}
}

func (c Service) GetSession(r *http.Request) (map[string]interface{}, interfaces.ErrorWrapper) {
  cookie, err := r.Cookie(cookieSessionKey)
  if err != nil {
    return nil, models.NewErrorWrapper(err, NoAuthCookie)
  }

  decoded, err := c.coder.Decrypt(cookie.Value)
  if err != nil {
    return nil, models.NewErrorWrapper(err, InvalidAuthCookie)
  }

  return decoded, nil
}

func (c Service) SetSession(r *http.Request, session map[string]interface{}) interfaces.ErrorWrapper {
  token, err := c.coder.Encrypt(session)
  if err != nil {
    return models.NewErrorWrapper(err, services.InternalError)
  }

  r.AddCookie(&http.Cookie{
    Name: cookieSessionKey,
    Value: token,
    Path: "/",
    HttpOnly: true,
    MaxAge: 3600,
  })
  return nil
}

func (c Service) InvalidateSession(r *http.Request) {
  r.AddCookie(&http.Cookie{
    Name: cookieSessionKey,
    Value: "",
    Path: "/",
    HttpOnly: true,
    Expires: time.Unix(0, 0),
  })
}
