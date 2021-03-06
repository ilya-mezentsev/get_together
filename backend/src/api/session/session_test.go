package session

import (
	"api"
	"api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	sessionAPIMock "mock/api"
	repositoriesMock "mock/repositories"
	"models"
	"os"
	"plugins/config"
	"repositories"
	"services"
	"services/errors"
	"services/proxies/validation"
	"testing"
	"utils"
)

var (
	db     *sqlx.DB
	router = api.GetRouter()
)

func init() {
	utils.SkipInShortMode()

	coderKey, err := config.GetCoderKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err = config.GetConfiguredConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sessionService := services.Session(coderKey)
	InitRequestHandlers(
		services.Authentication(repositories.Credentials(db)),
		sessionService,
		middlewares.AuthSession{Service: sessionService}.HasValidSession,
	)
}

func TestMain(m *testing.M) {
	repositoriesMock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestSessionGet_Success(t *testing.T) {
	var response sessionAPIMock.GetSessionSuccess
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.RequestWithSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(float64(sessionAPIMock.TestSessionData.Id), response.Data["id"], t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionGet_NoSessionError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.RequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.NoAuthCookie.Error(), response.ErrorDetail, t)
}

func TestSessionGet_InvalidSessionError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.RequestWithInvalidToken(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InvalidAuthCookie.Error(), response.ErrorDetail, t)
}

func TestSessionRegister_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessRegistrationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionRegister_EmailExistsError(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.EmailExistsRegistrationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.EmailExists.Error(), response.ErrorDetail, t)
}

func TestSessionRegister_InvalidEmailError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidEmailRegistrationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidEmail, response.ErrorDetail, t)
}

func TestSessionRegister_InvalidPasswordError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidPasswordRegistrationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidPassword, response.ErrorDetail, t)
}

func TestSessionRegister_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessRegistrationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionLogin_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessLoginRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionLogin_CredentialsNotFoundError(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.NoCredentialsLoginRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.CredentialsNotFound.Error(), response.ErrorDetail, t)
}

func TestSessionLogin_InvalidEmailError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidEmailLoginRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidEmail, response.ErrorDetail, t)
}

func TestSessionLogin_InvalidPasswordError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidPasswordLoginRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidPassword, response.ErrorDetail, t)
}

func TestSessionLogin_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessLoginRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionChangePassword_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessChangePasswordRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionChangePassword_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessChangePasswordRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestSessionChangePassword_UserIdNotFoundError(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.UserIdNotFoundChangePasswordRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestSessionChangePassword_InvalidUserIdError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidUserIdChangePasswordRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestSessionChangePassword_InvalidPasswordError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.InvalidPasswordChangePasswordRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidPassword, response.ErrorDetail, t)
}

func TestSessionChangePassword_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessChangePasswordRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionLogout_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessLogoutRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionLogout_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(sessionAPIMock.SuccessLogoutRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}
