package authentication

import (
  "interfaces"
  "internal_errors"
  "models"
  "plugins/logger"
  "services"
  "utils"
)

type AuthService struct {
  repository interfaces.CredentialsRepository
}

func New(repository interfaces.CredentialsRepository) AuthService {
  return AuthService{repository: repository}
}

func (s AuthService) RegisterUser(credentials models.UserCredentials) error {
  credentials.Password = s.generatePassword(credentials)

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

func (s AuthService) generatePassword(credentials models.UserCredentials) string {
  return utils.GetHash(credentials.Email + credentials.Password)
}

func (s AuthService) Login(credentials models.UserCredentials) (uint, error) {
  credentials.Password = s.generatePassword(credentials)
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

func (s AuthService) ChangePassword(userId uint, password string) error {
  email, err := s.getUserEmail(userId)
  if err != nil {
    return err
  }

  user := models.UserCredentials{Email: email, Password: password}
  user.Password = s.generatePassword(user)
  switch err := s.repository.UpdateUserPassword(user); err {
  case nil:
    return nil
  default:
    logger.ErrorF("unable to change user password: %v", err)
    return services.InternalError
  }
}

func (s AuthService) getUserEmail(userId uint) (string, error) {
  email, err := s.repository.GetUserEmail(userId)

  switch err {
  case nil:
    return email, nil
  case internal_errors.UnableToFindUserById:
    logger.Warning(err)
    return "", services.UserIdNotFound
  default:
    logger.ErrorF("unable to get user email: %v", err)
    return "", services.InternalError
  }
}
