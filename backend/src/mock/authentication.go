package mock

import (
  "errors"
  "internal_errors"
  "models"
  "strings"
)

type CredentialsMock struct {
  Users []models.UserCredentials
}

var (
  Users = []models.UserCredentials{
    {Email: "mail@ya.ru", Password: "hello_world"},
    {Email: "mail@gmail.com", Password: "hi_there"},
  }
  CredentialsRepo = CredentialsMock{
    Users: Users,
  }

  NewUser = models.UserCredentials{
    Email: "test@test.ru",
    Password: "pass",
  }
  ExistingUserEmail = models.UserCredentials{
    Email: "mail@ya.ru",
    Password: "pass",
  }
  WrongUserPassword = models.UserCredentials{
    Email: "test@ya.ru",
    Password: strings.Repeat("big string", 100),
  }

  tooLongPasswordError = errors.New("too long password")
)

func (c *CredentialsMock) ResetState() {
  c.Users = Users
}

func (c *CredentialsMock) CreateUser(user models.UserCredentials) error {
  if c.HasEmail(user.Email) {
    return internal_errors.UnableToRegisterUserEmailExists
  } else if len(user.Password) > 32 {
    return tooLongPasswordError
  }

  c.Users = append(c.Users, user)
  return nil
}

func (c *CredentialsMock) HasEmail(email string) bool {
  for _, user := range c.Users {
    if user.Email == email {
      return true
    }
  }

  return false
}

func (c *CredentialsMock) GetUserIdByCredentials(user models.UserCredentials) (uint, error) {
  if len(user.Password) > 32 {
    return 0, tooLongPasswordError
  }

  for userIdx, registered := range c.Users {
    if registered == user {
      return uint(userIdx)+1, nil
    }
  }

  return 0, internal_errors.UnableToLoginUserNotFound
}

