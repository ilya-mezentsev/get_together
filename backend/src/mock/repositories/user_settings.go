package repositories

import (
  "github.com/jmoiron/sqlx"
  "models"
)

var (
  UsersInfoQueries = []string{
    "INSERT INTO users_info(user_id, name, nickname, gender, age) VALUES(1, 'J. Smith', 'mather_fucker', 'male', 12)",
    `INSERT INTO users_info(user_id, name, nickname, age, avatar_url) VALUES(2, 'Mr. Anderson', 'LoL228', 8, 'http://123.png')`,
  }
  UsersRatingQueries = []string{
    "INSERT INTO users_rating(user_id, tag, value) VALUES(1, 'tag1', 85)",
    "INSERT INTO users_rating(user_id, tag, value) VALUES(1, 'tag2', 75)",
    "INSERT INTO users_rating(user_id, tag, value) VALUES(2, 'tag1', 95)",
    "INSERT INTO users_rating(user_id, tag, value) VALUES(2, 'tag3', 65)",
  }

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

func InitUsersSettings(db *sqlx.DB) {
  dropTables(db)
  initTables(db)

  for _, q := range append(UsersQueries, append(UsersInfoQueries, UsersRatingQueries...)...) {
    _, err := db.Exec(q)
    if err != nil {
      panic(err)
    }
  }
}

func DropUsersSettings(db *sqlx.DB) {
  dropTables(db)
}
