package repositories

import (
  "github.com/jmoiron/sqlx"
  "models"
)

var (
  UsersQueries = []string{
    "INSERT INTO users(email, password) VALUES('mail@ya.ru', '3dac4de4c9d5af7382da4c63f5555f2b')",
    "INSERT INTO users(email, password) VALUES('me@gmail.com', '0c120226ef10689396a6eabbf733e54b')",
  }
  Users = []models.UserCredentials{
    { Email: "mail@ya.ru", Password: "3dac4de4c9d5af7382da4c63f5555f2b" },
    { Email: "me@gmail.com", Password: "0c120226ef10689396a6eabbf733e54b" },
  }
  NotExistsUser = models.UserCredentials{
    Email: "no-way@ya.ru",
    Password: "",
  }
  NewUser =  models.UserCredentials{
    Email: "new@mail.ru",
    Password: "some_pass",
  }
)

func InitUsers(db *sqlx.DB) {
  dropTables(db)
  initTables(db)

  for _, q := range UsersQueries {
    _, err := db.Exec(q)
    if err != nil {
      panic(err)
    }
  }
}

func DropUsers(db *sqlx.DB) {
  dropTables(db)
}
