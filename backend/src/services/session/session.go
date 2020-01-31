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
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to decrypt session: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "session": cookie.Value,
      },
    }, logger.Warning)
    return nil, InvalidAuthCookie
  }

  return decoded, nil
}

func (c Controller) SetSession(r *http.Request, session map[string]interface{}) error {
  token, err := c.coder.Encrypt(session)
  if err != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to encrypt session: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "session": session,
      },
    }, logger.Warning)
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
