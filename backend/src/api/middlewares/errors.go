package middlewares

import (
	"api"
	"errors"
)

var (
	NoSession         = api.ApplicationError{OriginalError: errors.New("no-session")}
	NoCSRFCookie      = api.ApplicationError{OriginalError: errors.New("no-csrf-cookie")}
	InvalidCSRFCookie = api.ApplicationError{OriginalError: errors.New("invalid-csrf-cookie")}
	NoCSRFHeader      = api.ApplicationError{OriginalError: errors.New("no-csrf-header")}
	InvalidCSRFHeader = api.ApplicationError{OriginalError: errors.New("invalid-csrf-header")}
	InvalidCSRFToken  = api.ApplicationError{OriginalError: errors.New("invalid-csrf-token")}
	CSRFInternalError = api.ApplicationError{OriginalError: errors.New("csrf-internal-error")}
)
