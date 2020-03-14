package api

import "errors"

var (
	ReadRequestBodyError = errors.New("read request body error")
	CannotDecodeRequestBody = errors.New("unable to decode request body")
	CannotWriteResponse = errors.New("unable to write response")
)
