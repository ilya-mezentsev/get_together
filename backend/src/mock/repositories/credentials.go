package repositories

import (
  "models"
)

var (
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
