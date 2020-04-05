package middlewares

import (
	"api"
	"errors"
)

var (
	NoSession = api.ApplicationError{OriginalError: errors.New("no-session")}
)
