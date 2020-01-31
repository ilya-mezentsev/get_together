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
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "credentials": credentials,
      },
    }, logger.Warning)
    return EmailExists
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to register user: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "credentials": credentials,
      },
    }, logger.Error)
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
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "credentials": credentials,
      },
    }, logger.Warning)
    return 0, CredentialsNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to login user: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "credentials": credentials,
      },
    }, logger.Error)
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
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to change user password: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Error)
    return services.InternalError
  }
}

func (s AuthService) getUserEmail(userId uint) (string, error) {
  email, err := s.repository.GetUserEmail(userId)

  switch err {
  case nil:
    return email, nil
  case internal_errors.UnableToFindUserById:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Warning)
    return "", services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to get user email: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Error)
    return "", services.InternalError
  }
}
