package authentication

import "errors"

var (
  EmailExists = errors.New("email-exists")
  CredentialsNotFound = errors.New("credentials-not-found")
)
