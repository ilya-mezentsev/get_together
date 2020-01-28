package repositories

import (
  "github.com/jmoiron/sqlx"
  "models"
  "strings"
)

var (
  UsersQueries = []string{
    "INSERT INTO users(email, password) VALUES('mail@ya.ru', 'pass')",
    "INSERT INTO users(email, password) VALUES('me@gmail.com', 'root')",
  }
  Users = []models.UserCredentials{
    { Email: "mail@ya.ru", Password: "pass" },
    { Email: "me@gmail.com", Password: "root" },
  }
  NotExistsUser = models.UserCredentials{
    Email: "no-way@ya.ru",
    Password: "",
  }
  NewUser =  models.UserCredentials{
    Email: "new@mail.ru",
    Password: "some_pass",
  }
  IncorrectPasswordUser = models.UserCredentials{
    Email: "test@gmail.com",
    Password: strings.Repeat("long pass", 100),
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
