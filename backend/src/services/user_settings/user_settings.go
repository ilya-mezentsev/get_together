package user_settings

import (
  "interfaces"
  "internal_errors"
  "models"
  "services"
)

type Service struct {
  repository interfaces.UsersSettingsRepository
}

func New(repository interfaces.UsersSettingsRepository) Service {
  return Service{repository}
}

func (u Service) GetUserSettings(userId uint) (models.FullUserInfo, error) {
  info, err := u.repository.GetUserSettings(userId)

  switch err {
  case nil:
    return info, nil
  case internal_errors.UnableToFindUserById:
    return models.FullUserInfo{}, services.UserIdNotFound
  default:
    return models.FullUserInfo{}, services.InternalError
  }
}

func (u Service) UpdateUserSettings(userId uint, info models.UserSettings) error {
  switch err := u.repository.UpdateUserSettings(userId, info); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    return services.UserIdNotFound
  default:
    return services.InternalError
  }
}
