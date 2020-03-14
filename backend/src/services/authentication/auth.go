package authentication

import (
  "interfaces"
  "internal_errors"
  "models"
  "services/errors"
  "utils"
)

type Service struct {
  repository interfaces.CredentialsRepository
}

func New(repository interfaces.CredentialsRepository) Service {
  return Service{repository}
}

func (s Service) RegisterUser(credentials models.UserCredentials) error {
  credentials.Password = s.generatePassword(credentials)

  switch err := s.repository.CreateUser(credentials); err {
  case nil:
    return nil
  case internal_errors.UnableToRegisterUserEmailExists:
    return errors.EmailExists
  default:
    return errors.InternalError
  }
}

func (s Service) generatePassword(credentials models.UserCredentials) string {
  return utils.GetHash(credentials.Email + credentials.Password)
}

func (s Service) Login(credentials models.UserCredentials) (models.UserSession, error) {
  credentials.Password = s.generatePassword(credentials)
  userId, err := s.repository.GetUserIdByCredentials(credentials)

  switch err {
  case nil:
    return models.UserSession{ID: userId}, nil
  case internal_errors.UnableToLoginUserNotFound:
    return models.UserSession{}, errors.CredentialsNotFound
  default:
    return models.UserSession{}, errors.InternalError
  }
}

func (s Service) ChangePassword(userId uint, password string) error {
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
    return errors.InternalError
  }
}

func (s Service) getUserEmail(userId uint) (string, error) {
  email, err := s.repository.GetUserEmail(userId)

  switch err {
  case nil:
    return email, nil
  case internal_errors.UnableToFindUserById:
    return "", errors.UserIdNotFound
  default:
    return "", errors.InternalError
  }
}
