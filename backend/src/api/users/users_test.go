package users

import (
	"api"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	usersAPIMock "mock/api"
	mock "mock/repositories"
	"models"
	"os"
	"repositories"
	"services"
	"services/errors"
	"services/proxies/validation"
	"strings"
	"testing"
	"utils"
)

var (
	db *sqlx.DB
	router = api.GetRouter()
)

func init() {
	utils.SkipInShortMode()

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

	InitRequestHandlers(services.UserSettings(repositories.UserSettings(db)))
}

func TestMain(m *testing.M) {
	mock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestUserSettingsGet_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response usersAPIMock.UserSettingsResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.FirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertEqual(mock.UsersInfo[0]["name"], response.Data.Name, t)
	utils.AssertEqual(mock.UsersInfo[0]["nickname"], response.Data.Nickname, t)
	utils.AssertEqual(mock.UsersInfo[0]["gender"], response.Data.Gender, t)
}

func TestUserSettingsGet_UserIdNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.NotExistsUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestUserSettingsGet_InvalidIdError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.InvalidIDUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestUserSettingsGet_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.FirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestUserSettingsPatch_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchFirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestUserSettingsPatch_UserIdNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchNotExistsUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestUserSettingsPatch_InvalidUserIdError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.InvalidIdUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestUserSettingsPatch_InvalidAllSettingsUserIdError(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.InvalidAllSettingsUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	for _, validationError := range []string{
		validation.InvalidUserName, validation.InvalidUserNickname,
		validation.InvalidUserAge, validation.InvalidUserAvatarURL,
	} {
		utils.AssertTrue(strings.Contains(response.ErrorDetail, validationError), t)
	}
}

func TestUserSettingsPatch_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchNotExistsUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}
