package users

import (
	"api"
	"api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"interfaces"
	"io/ioutil"
	"log"
	usersAPIMock "mock/api"
	repositoriesMock "mock/repositories"
	"models"
	"os"
	"plugins/config"
	"repositories"
	"services"
	"services/errors"
	"services/proxies/validation"
	"strings"
	"testing"
	"utils"
)

var (
	db             *sqlx.DB
	sessionService interfaces.SessionAccessorService
	router         = api.GetRouter()
)

func init() {
	utils.SkipInShortMode()

	var err error
	db, err = config.GetConfiguredConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	coderKey, err := config.GetCoderKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sessionService = services.Session(coderKey)
	InitRequestHandlers(
		services.UserSettings(repositories.UserSettings(db)),
		middlewares.AuthSession{Service: sessionService}.HasValidSession,
	)
}

func TestMain(m *testing.M) {
	repositoriesMock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestUserSettingsGet_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response usersAPIMock.UserSettingsResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.FirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertEqual(repositoriesMock.UsersInfo[0]["name"], response.Data.Name, t)
	utils.AssertEqual(repositoriesMock.UsersInfo[0]["nickname"], response.Data.Nickname, t)
	utils.AssertEqual(repositoriesMock.UsersInfo[0]["gender"], response.Data.Gender, t)
}

func TestUserSettingsGet_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.FirstUserSettingsRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestUserSettingsGet_UserIdNotFoundError(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

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
		utils.MakeRequest(usersAPIMock.InvalidIdUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestUserSettingsGet_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.FirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestUserSettingsPatch_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchFirstUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestUserSettingsPatch_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchFirstUserSettingsRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestUserSettingsPatch_UserIdNotFoundError(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

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
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(usersAPIMock.PatchNotExistsUserSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}
