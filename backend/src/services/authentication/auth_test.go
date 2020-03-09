package authentication

import (
  mock "mock/services"
  "services"
  "testing"
  "utils"
)

var authService = New(&mock.CredentialsRepo)

func TestAuthService_RegisterUserSuccess(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.NewUser)
  utils.AssertNil(err, t)

  userSession, err := authService.Login(mock.NewUser)
  utils.AssertNil(err, t)
  utils.AssertEqual(int(userSession.ID), len(mock.CredentialsRepo.Users), t)
}

func TestAuthService_RegisterUserEmailExistsError(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.ExistingUserEmail)
  utils.AssertErrorsEqual(EmailExists, err, t)
}

func TestAuthService_RegisterUserInternalError(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.RegisterUser(mock.BadUser)
  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestAuthService_LoginSuccess(t *testing.T) {
  userSession, err := authService.Login(mock.Users[0])

  utils.AssertNil(err, t)
  utils.AssertEqual(1, int(userSession.ID), t)
}

func TestAuthService_LoginCredentialsNotFoundError(t *testing.T) {
  _, err := authService.Login(mock.NewUser)

  utils.AssertErrorsEqual(CredentialsNotFound, err, t)
}

func TestAuthService_LoginInternalError(t *testing.T) {
  _, err := authService.Login(mock.BadUser)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestAuthService_ChangePasswordSuccess(t *testing.T) {
  defer mock.CredentialsRepo.ResetState()

  err := authService.ChangePassword(1, "new_password")
  email, _ := mock.CredentialsRepo.GetUserEmail(1)
  expectedPassword := utils.GetHash(email + "new_password")

  utils.AssertNil(err, t)
  utils.AssertEqual(expectedPassword, mock.CredentialsRepo.Users[0].Password, t)
}

func TestAuthService_ChangePasswordUserNotFoundError(t *testing.T) {
  err := authService.ChangePassword(11, "")

  utils.AssertErrorsEqual(services.UserIdNotFound, err, t)
}

func TestAuthService_ChangePasswordInternalError_1(t *testing.T) {
  err := authService.ChangePassword(1, mock.BadUser.Password)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestAuthService_ChangePasswordInternalError_2(t *testing.T) {
  err := authService.ChangePassword(mock.BadUserId, "")

  utils.AssertErrorsEqual(services.InternalError, err, t)
}
