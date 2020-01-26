package session

import "errors"

var (
  NoAuthCookie = errors.New("no-auth-cookie")
  InvalidAuthCookie = errors.New("invalid-auth-cookie")
)
