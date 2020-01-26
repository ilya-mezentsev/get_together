package internal_errors

import "errors"

var (
  UnableToRegisterUserEmailExists = errors.New("unable to register user: email already exists")
  UnableToLoginUserNotFound = errors.New("unable to login user: not found by credentials")
)
