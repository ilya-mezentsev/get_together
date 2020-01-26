package session

import (
  "net/http"
  "plugins/code"
  "plugins/logger"
  "services"
  "time"
)

const cookieSessionKey = "GT-Session-Token"

type Controller struct {
  coder code.Coder
}

func New(key string) Controller {
  return Controller{coder: code.NewCoder(key)}
}

func (c Controller) GetSession(r *http.Request) (map[string]interface{}, error) {
  cookie, err := r.Cookie(cookieSessionKey)
  if err != nil {
    logger.WarningF("unable to get session from cookie: %v", err)
    return nil, NoAuthCookie
  }

  decoded, err := c.coder.Decrypt(cookie.Value)
  if err != nil {
    logger.WarningF("unable to decode cookie '%s', error: %v", cookie.Value, err)
    return nil, InvalidAuthCookie
  }

  return decoded, nil
}

func (c Controller) SetSession(r *http.Request, session map[string]interface{}) error {
  token, err := c.coder.Encrypt(session)
  if err != nil {
    logger.WarningF("unable to encrypt session %v, error: %v", session, err)
    return services.InternalError
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

func (c Controller) InvalidateSession(r *http.Request) {
  r.AddCookie(&http.Cookie{
    Name: cookieSessionKey,
    Value: "",
    Path: "/",
    HttpOnly: true,
    Expires: time.Unix(0, 0),
  })
}
