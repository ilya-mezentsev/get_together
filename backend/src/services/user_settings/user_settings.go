package user_settings

import (
	"interfaces"
	"internal_errors"
	"models"
	"services/errors"
)

type Service struct {
  repository interfaces.UsersSettings
}

func New(repository interfaces.UsersSettings) Service {
  return Service{repository}
}

func (u Service) GetUserSettings(userId uint) (models.FullUserInfo, error) {
  info, err := u.repository.GetUserSettings(userId)

  switch err {
  case nil:
    return info, nil
  case internal_errors.UnableToFindUserById:
    return models.FullUserInfo{}, errors.UserIdNotFound
  default:
    return models.FullUserInfo{}, errors.InternalError
  }
}

func (u Service) UpdateUserSettings(userId uint, info models.UserSettings) error {
  switch err := u.repository.UpdateUserSettings(userId, info); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    return errors.UserIdNotFound
  default:
    return errors.InternalError
  }
}
