package session

import (
  "api"
  "encoding/json"
  "fmt"
  "github.com/jmoiron/sqlx"
  "io/ioutil"
  "log"
  mock "mock/repositories"
  sessionMock "mock/services"
  "models"
  "os"
  "repositories"
  "services"
  "services/authentication"
  sessionService "services/session"
  "testing"
  "utils"
)

var (
  db *sqlx.DB
  router = api.GetRouter()
)

func init() {
  utils.SkipInShortMode()

  coderKey := os.Getenv("CODER_KEY")
  if coderKey == "" {
    fmt.Println("CODER_KEY env var is not set")
    os.Exit(1)
  }

  connStr := os.Getenv("CONN_STR")
  if connStr == "" {
    fmt.Println("CONN_STR env var is not set")
    os.Exit(1)
  }

  var err error
  db, err = sqlx.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  InitRequestHandlers(
    authentication.New(repositories.Credentials(db)), sessionService.New(coderKey))
}

func TestMain(m *testing.M) {
  mock.DropTables(db)
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestSessionGet_Success(t *testing.T) {
  var response sessionMock.GetSessionSuccess
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.RequestWithSession(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(float64(sessionMock.TestSessionData.ID), response.Data["id"], t)
  utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionGet_NoSessionError(t *testing.T) {
  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.RequestWithoutSession(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(sessionService.NoAuthCookie.Error(), response.ErrorDetail, t)
}

func TestSessionGet_InvalidSessionError(t *testing.T) {
  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.RequestWithInvalidToken(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(sessionService.InvalidAuthCookie.Error(), response.ErrorDetail, t)
}

func TestSessionRegister_Success(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response sessionMock.DefaultSuccess
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessRegistrationRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestSessionRegister_EmailExistsError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.EmailExistsRegistrationRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(authentication.EmailExists.Error(), response.ErrorDetail, t)
}

func TestSessionRegister_InternalError(t *testing.T) {
  mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessRegistrationRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(services.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionLogin_Success(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.SuccessResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessLoginRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusOk, response.Status, t)
  utils.AssertNil(response.Data, t)
}

func TestSessionLogin_CredentialsNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.NoCredentialsLoginRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(authentication.CredentialsNotFound.Error(), response.ErrorDetail, t)
}

func TestSessionLogin_InternalError(t *testing.T) {
  mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessLoginRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(services.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionChangePassword_Success(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.SuccessResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessChangePasswordRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusOk, response.Status, t)
  utils.AssertNil(response.Data, t)
}

func TestSessionChangePassword_UserIdNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.UserIdNotFoundChangePasswordRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(services.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestSessionChangePassword_InternalError(t *testing.T) {
  mock.DropTables(db)

  var response models.ErrorResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessChangePasswordRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusError, response.Status, t)
  utils.AssertEqual(services.InternalError.Error(), response.ErrorDetail, t)
}

func TestSessionLogout_Success(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  var response models.SuccessResponse
  err := json.NewDecoder(
    utils.MakeRequest(sessionMock.SuccessLogoutRequest(router))).Decode(&response)

  utils.AssertNil(err, t)
  utils.AssertEqual(api.StatusOk, response.Status, t)
  utils.AssertNil(response.Data, t)
}
