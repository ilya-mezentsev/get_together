package chats

import (
	"api"
	"api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"interfaces"
	"io/ioutil"
	"log"
	chatAPIMock "mock/api"
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
	chatRepository := repositories.Chat(db)
	InitRequestHandlers(
		services.Chat(chatRepository),
		services.ChatAccessor(chatRepository),
		middlewares.AuthSession{Service: sessionService}.HasValidSession,
	)
}

func TestMain(m *testing.M) {
	repositoriesMock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetMeetingChat_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response chatAPIMock.MeetingChatResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertEqual(repositoriesMock.MeetingType, response.Data.Type, t)
}

func TestGetMeetingChat_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetMeetingChatRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestGetMeetingChat_ChatNotFound(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetMeetingChatByMeetingIfWithoutChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestGetMeetingChat_InvalidMeetingId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetMeetingChatInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestGetMeetingChat_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestGetUserChats_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response chatAPIMock.UserChatsResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetUserChatsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNotNil(response.Data, t)
}

func TestGetUserChats_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetUserChatsWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestGetUserChats_InvalidId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetUserChatsInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestGetUserChats_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.GetUserChatsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestCreateMeetingChat_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestCreateMeetingChat_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingChatRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestCreateMeetingChat_InvalidId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingChatInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestCreateMeetingChat_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestCreateMeetingRequestChat_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingRequestChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestCreateMeetingRequestChat_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingRequestChatRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestCreateMeetingRequestChat_InvalidId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingRequestChatInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestCreateMeetingRequestChat_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CreateMeetingRequestChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingChat_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestCloseMeetingChat_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingChatRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingChat_IdNotFound(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingChatIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.ChatIdNotFound.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingChat_InvalidId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingChatInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestCloseMeetingChat_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingRequestChat_Success(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingRequestChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestCloseMeetingRequestChat_NoSession(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingRequestChatRequestWithoutSession(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(middlewares.NoSession.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingRequestChat_IdNotFound(t *testing.T) {
	repositoriesMock.InitTables(db)
	defer repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingRequestChatIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.ChatIdNotFound.Error(), response.ErrorDetail, t)
}

func TestCloseMeetingRequestChat_InvalidId(t *testing.T) {
	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingRequestChatInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
}

func TestCloseMeetingRequestChat_InternalError(t *testing.T) {
	repositoriesMock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(chatAPIMock.CloseMeetingRequestChatRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}
