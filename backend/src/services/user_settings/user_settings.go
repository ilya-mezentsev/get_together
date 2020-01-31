package user_settings

import (
  "interfaces"
  "internal_errors"
  "models"
  "plugins/logger"
  "services"
)

type UsersSettingsService struct {
  repository interfaces.UsersSettingsRepository
}

func New(repository interfaces.UsersSettingsRepository) UsersSettingsService {
  return UsersSettingsService{repository: repository}
}

func (u UsersSettingsService) GetUserSettings(userId uint) (models.FullUserInfo, error) {
  info, err := u.repository.GetUserSettings(userId)

  switch err {
  case nil:
    return info, nil
  case internal_errors.UnableToFindUserById:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Warning)
    return models.FullUserInfo{}, services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to get user info: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Error)
    return models.FullUserInfo{}, services.InternalError
  }
}

func (u UsersSettingsService) UpdateUserSettings(userId uint, info models.UserSettings) error {
  switch err := u.repository.UpdateUserSettings(userId, info); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "user_id": userId,
        "info": info,
      },
    }, logger.Warning)
    return services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to update user info: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "user_id": userId,
        "info": info,
      },
    }, logger.Error)
    return services.InternalError
  }
}
