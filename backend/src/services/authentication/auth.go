package authentication

import (
  "interfaces"
  "internal_errors"
  "models"
  "plugins/logger"
  "services"
)

type AuthService struct {
  repository interfaces.CredentialsRepository
}

func New(repository interfaces.CredentialsRepository) AuthService {
  return AuthService{repository: repository}
}

func (s AuthService) RegisterUser(credentials models.UserCredentials) error {
  switch err := s.repository.CreateUser(credentials); err {
  case nil:
    return nil
  case internal_errors.UnableToRegisterUserEmailExists:
    logger.Warning(err)
    return EmailExists
  default:
    logger.ErrorF("unable to register user: %v", err)
    return services.InternalError
  }
}

func (s AuthService) Login(credentials models.UserCredentials) (uint, error) {
  userId, err := s.repository.GetUserIdByCredentials(credentials)

  switch err {
  case nil:
    return userId, nil
  case internal_errors.UnableToLoginUserNotFound:
    logger.Warning(err)
    return 0, CredentialsNotFound
  default:
    logger.ErrorF("unable to login user: %v", err)
    return 0, services.InternalError
  }
}
