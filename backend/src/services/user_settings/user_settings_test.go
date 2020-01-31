package user_settings

import (
  "io/ioutil"
  "log"
  mock "mock/services"
  "os"
  "services"
  "testing"
  "utils"
)

var service = New(&mock.UsersSettingsRepository)

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestUsersSettingsService_GetUserInfoSuccess(t *testing.T) {
  defer mock.UsersSettingsRepository.ResetState()

  full, err := service.GetUserSettings(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.UsersSettingsRepository.Settings[1].UserSettings == full.UserSettings, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.UsersSettingsRepository.Settings[1].UserSettings, Got: full.UserSettings}))
    t.Fail()
  })
}

func TestUsersSettingsService_GetUserInfoUserNotFoundError(t *testing.T) {
  _, err := service.GetUserSettings(11)

  utils.AssertErrorsEqual(services.UserIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestUsersSettingsService_GetUserInfoInternalError(t *testing.T) {
  _, err := service.GetUserSettings(mock.BadUserId)

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestUsersSettingsService_UpdateUserInfoSuccess(t *testing.T) {
  defer mock.UsersSettingsRepository.ResetState()

  err := service.UpdateUserSettings(1, mock.NewUserInfo)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.UsersSettingsRepository.Settings[1].UserSettings == mock.NewUserInfo, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.UsersSettingsRepository.Settings[1].UserSettings, Got: mock.NewUserInfo}))
    t.Fail()
  })
}

func TestUsersSettingsService_UpdateUserInfoUserNotFoundError(t *testing.T) {
  err := service.UpdateUserSettings(11, mock.NewUserInfo)

  utils.AssertErrorsEqual(services.UserIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestUsersSettingsService_UpdateUserInfoInternalError(t *testing.T) {
  err := service.UpdateUserSettings(mock.BadUserId, mock.NewUserInfo)

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}
