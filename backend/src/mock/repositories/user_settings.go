package repositories

import (
  "models"
)

var (
  FirstUserSettings = models.FullUserInfo{
    UserSettings: models.UserSettings{
      Name: "J. Smith", Nickname: "mather_fucker", Gender: "male", Age: 12,
    },
    Rating: []models.Rating{
      {Tag: "tag1", Value: 85},
      {Tag: "tag2", Value: 75},
    },
  }
  TestInfo = models.UserSettings{
    Nickname: "some_nickname",
    Gender: "male",
  }
)

func SettingsEqual(s1, s2 models.FullUserInfo) bool {
  if s1.UserSettings != s2.UserSettings {
    return false
  }

  if len(s1.Rating) != len(s2.Rating) {
    return false
  }

  for i := range s1.Rating {
    if s1.Rating[i] != s2.Rating[i] {
      return false
    }
  }

  return true
}
