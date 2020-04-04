package api

import "errors"

type ApplicationError struct {
	OriginalError error
}

var (
	ReadRequestBodyError    = ApplicationError{errors.New("read-request-body-error")}
	CannotDecodeRequestBody = ApplicationError{errors.New("decode-request-body-error")}
)

func (a ApplicationError) Error() string {
	return a.OriginalError.Error()
}
