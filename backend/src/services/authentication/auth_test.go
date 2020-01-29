package authentication

import (
  "io/ioutil"
  "log"
  mock "mock/services"
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

  userId, err := authService.Login(mock.NewUser)
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

func TestAuthService_RegisterUserInternalError(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.BadUser)
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

func TestAuthService_LoginInternalError(t *testing.T) {
  _, err := authService.Login(mock.BadUser)

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_ChangePasswordSuccess(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.ChangePassword(1, "new_password")
  email, _ := mock.CredentialsRepo.GetUserEmail(1)
  expectedPassword := utils.GetHash(email + "new_password")

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(expectedPassword == mock.CredentialsRepo.Users[0].Password, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expectedPassword, Got: mock.CredentialsRepo.Users[0].Password}))
    t.Fail()
  })
}

func TestAuthService_ChangePasswordUserNotFoundError(t *testing.T) {
  err := authService.ChangePassword(11, "")

  utils.AssertErrorsEqual(services.UserIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_ChangePasswordInternalError_1(t *testing.T) {
  err := authService.ChangePassword(1, mock.BadUser.Password)

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestAuthService_ChangePasswordInternalError_2(t *testing.T) {
  err := authService.ChangePassword(mock.BadUserId, "")

  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}
