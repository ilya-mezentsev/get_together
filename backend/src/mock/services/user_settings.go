package services

import (
  "internal_errors"
  "models"
)

type UsersSettingsRepositoryMock struct {
  Settings map[uint]models.FullUserInfo
}

var (
  usersSettings = map[uint]models.FullUserInfo{
    1: {
      UserSettings: models.UserSettings{
        Name: "J. Smith",
        Age: 15,
      },
      Rating: nil,
    },
    2: {
      UserSettings:   models.UserSettings{
        Name: "Nick",
        Age: 16,
      },
      Rating: nil,
    },
  }
  UsersSettingsRepository = UsersSettingsRepositoryMock{Settings: usersSettings}
  NewUserInfo = models.UserSettings{
    Name: "Hello world",
    Age: 16,
    Gender: "male",
  }
)

func (u *UsersSettingsRepositoryMock) ResetState() {
  u.Settings = usersSettings
}

func (u *UsersSettingsRepositoryMock) GetUserSettings(userId uint) (models.FullUserInfo, error) {
  if userId == BadUserId {
    return models.FullUserInfo{}, someInternalError
  }

  userInfo, found := u.Settings[userId]
  if !found {
    return models.FullUserInfo{}, internal_errors.UnableToFindUserById
  }

  return userInfo, nil
}

func (u *UsersSettingsRepositoryMock) UpdateUserSettings(userId uint, info models.UserSettings) error {
  if userId == BadUserId {
    return someInternalError
  }

  _, found := u.Settings[userId]
  if !found {
    return internal_errors.UnableToFindUserById
  }

  u.Settings[userId] = models.FullUserInfo{
    UserSettings: info,
    Rating:       nil,
  }
  return nil
}

