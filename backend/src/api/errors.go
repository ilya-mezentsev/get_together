package api

import "errors"

var (
	ReadRequestBodyError    = errors.New("read-request-body-error")
	CannotDecodeRequestBody = errors.New("decode-request-body-error")
	CannotWriteResponse     = errors.New("write-response-error")
)
