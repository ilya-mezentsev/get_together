package authentication

import (
  "interfaces"
  "internal_errors"
  "models"
  "services"
  "utils"
)

type Service struct {
  repository interfaces.CredentialsRepository
}

func New(repository interfaces.CredentialsRepository) Service {
  return Service{repository: repository}
}

func (s Service) RegisterUser(credentials models.UserCredentials) interfaces.ErrorWrapper {
  credentials.Password = s.generatePassword(credentials)

  switch err := s.repository.CreateUser(credentials); err {
  case nil:
    return nil
  case internal_errors.UnableToRegisterUserEmailExists:
    return models.NewErrorWrapper(err, EmailExists)
  default:
    return models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) generatePassword(credentials models.UserCredentials) string {
  return utils.GetHash(credentials.Email + credentials.Password)
}

func (s Service) Login(credentials models.UserCredentials) (uint, interfaces.ErrorWrapper) {
  credentials.Password = s.generatePassword(credentials)
  userId, err := s.repository.GetUserIdByCredentials(credentials)

  switch err {
  case nil:
    return userId, nil
  case internal_errors.UnableToLoginUserNotFound:
    return 0, models.NewErrorWrapper(err, CredentialsNotFound)
  default:
    return 0, models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) ChangePassword(userId uint, password string) interfaces.ErrorWrapper {
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
    return models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) getUserEmail(userId uint) (string, interfaces.ErrorWrapper) {
  email, err := s.repository.GetUserEmail(userId)

  switch err {
  case nil:
    return email, nil
  case internal_errors.UnableToFindUserById:
    return "", models.NewErrorWrapper(err, services.UserIdNotFound)
  default:
    return "", models.NewErrorWrapper(err, services.InternalError)
  }
}
