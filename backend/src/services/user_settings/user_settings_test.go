package user_settings

import (
  mock "mock/services"
  "services"
  "testing"
  "utils"
)

var service = New(&mock.UsersSettingsRepository)

func TestUsersSettingsService_GetUserInfoSuccess(t *testing.T) {
  defer mock.UsersSettingsRepository.ResetState()

  full, err := service.GetUserSettings(1)
  utils.AssertNil(err, t)
  utils.AssertEqual(mock.UsersSettingsRepository.Settings[1].UserSettings, full.UserSettings, t)
}

func TestUsersSettingsService_GetUserInfoUserNotFoundError(t *testing.T) {
  _, err := service.GetUserSettings(11)

  utils.AssertErrorsEqual(services.UserIdNotFound, err, t)
}

func TestUsersSettingsService_GetUserInfoInternalError(t *testing.T) {
  _, err := service.GetUserSettings(mock.BadUserId)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestUsersSettingsService_UpdateUserInfoSuccess(t *testing.T) {
  defer mock.UsersSettingsRepository.ResetState()

  err := service.UpdateUserSettings(1, mock.NewUserInfo)
  utils.AssertNil(err, t)
  utils.AssertEqual(mock.UsersSettingsRepository.Settings[1].UserSettings, mock.NewUserInfo, t)
}

func TestUsersSettingsService_UpdateUserInfoUserNotFoundError(t *testing.T) {
  err := service.UpdateUserSettings(11, mock.NewUserInfo)

  utils.AssertErrorsEqual(services.UserIdNotFound, err, t)
}

func TestUsersSettingsService_UpdateUserInfoInternalError(t *testing.T) {
  err := service.UpdateUserSettings(mock.BadUserId, mock.NewUserInfo)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}
