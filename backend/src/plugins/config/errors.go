package config

import "errors"

var (
	noCoderKey         = errors.New("CODER_KEY env var is not set")
	noCSRFPrivateKey   = errors.New("CSRF_PRIVATE_KEY env var is not set")
	noConnectionString = errors.New("CONN_STR env var is not set")
	cannotOpenDB       = errors.New("cannot open DB")
)
