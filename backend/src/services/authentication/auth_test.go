package authentication

import (
	mock "mock/services"
	"services/errors"
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
	utils.AssertEqual(int(userSession.Id), len(mock.CredentialsRepo.Users), t)
}

func TestAuthService_RegisterUserEmailExistsError(t *testing.T) {
	defer mock.CredentialsRepo.ResetState()

	err := authService.RegisterUser(mock.FirstUserCredentials())
	utils.AssertErrorsEqual(errors.EmailExists, err, t)
}

func TestAuthService_RegisterUserInternalError(t *testing.T) {
	defer mock.CredentialsRepo.ResetState()

	err := authService.RegisterUser(mock.BadUser)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestAuthService_LoginSuccess(t *testing.T) {
	userSession, err := authService.Login(mock.FirstUserCredentials())

	utils.AssertNil(err, t)
	utils.AssertEqual(1, int(userSession.Id), t)
}

func TestAuthService_LoginCredentialsNotFoundError(t *testing.T) {
	_, err := authService.Login(mock.NewUser)

	utils.AssertErrorsEqual(errors.CredentialsNotFound, err, t)
}

func TestAuthService_LoginInternalError(t *testing.T) {
	_, err := authService.Login(mock.BadUser)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
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

	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestAuthService_ChangePasswordInternalError_1(t *testing.T) {
	err := authService.ChangePassword(1, mock.BadUser.Password)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestAuthService_ChangePasswordInternalError_2(t *testing.T) {
	err := authService.ChangePassword(mock.BadUserId, "")

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
