package messages

import "errors"

var (
	ReadJSONError = errors.New("unable-to-read-json-from-ws")
)
