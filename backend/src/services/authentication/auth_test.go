package authentication

import (
  "io/ioutil"
  "log"
  "mock"
  "os"
  "services"
  "testing"
  "utils"
)

var authService = New(&mock.CredentialsRepo)

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestAuthService_RegisterUserSuccess(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.NewUser)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })

  userId, err := mock.CredentialsRepo.GetUserIdByCredentials(mock.NewUser)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(int(userId) == len(mock.CredentialsRepo.Users), func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: len(mock.CredentialsRepo.Users), Got: userId}))
    t.Fail()
  })
}

func TestAuthService_RegisterUserEmailExistsError(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.ExistingUserEmail)
  utils.AssertErrorsEqual(EmailExists, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_RegisterUserWrongPasswordError(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.WrongUserPassword)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_LoginSuccess(t *testing.T) {
  userId, err := authService.Login(mock.Users[0])

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(1 == userId, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: 1, Got: userId}))
    t.Fail()
  })
}

func TestAuthService_LoginCredentialsNotFoundError(t *testing.T) {
  _, err := authService.Login(mock.NewUser)

  utils.AssertErrorsEqual(CredentialsNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_LoginWrongPasswordError(t *testing.T) {
  _, err := authService.Login(mock.WrongUserPassword)

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}
