package meetings

import (
	"api"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	meetingsAPIMock "mock/api"
	mock "mock/repositories"
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
	db     *sqlx.DB
	router = api.GetRouter()
)

func init() {
	utils.SkipInShortMode()
	var err error
	db, err = config.GetConfiguredConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	InitRequestHandlers(
		services.Meetings(repositories.Meetings(db)),
		services.Participation(repositories.UserSettings(db), repositories.MeetingsSettings(db)),
		services.MeetingsAccessor(repositories.Meetings(db)),
	)
}

func TestMain(m *testing.M) {
	mock.DropTables(db)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetPublicMeetings_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response meetingsAPIMock.PublicMeetingsResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetPublicMeetingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertEqual(len(mock.MeetingsSettings), len(response.Data), t)
}

func TestGetPublicMeetings_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetPublicMeetingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestGetExtendedMeetings_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response meetingsAPIMock.ExtendedMeetingsResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetExtendedMeetingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertEqual(len(mock.MeetingsSettings), len(response.Data), t)
}

func TestGetExtendedMeetings_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.GetExtendedMeetingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestCreateMeeting_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestCreateMeeting_NotExistsAdminID(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingNotExistsAdminIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestCreateMeeting_InvalidAdminID(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingInvalidAdminIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestCreateMeeting_InvalidMeetingSettings(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingWithInvalidSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	for _, validationError := range []string{
		validation.InvalidMeetingTitle, validation.InvalidMeetingDescription,
		validation.InvalidMeetingMaxUsers, validation.InvalidMeetingDuration,
		validation.InvalidMeetingMinAge, validation.InvalidMeetingLabel,
		validation.InvalidMeetingLatitude, validation.InvalidMeetingLongitude,
	} {
		utils.AssertTrue(strings.Contains(response.ErrorDetail, validationError), t)
	}
}

func TestDeleteMeeting_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestDeleteMeeting_MeetingIDNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingIDNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestDeleteMeeting_InvalidMeetingIDError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingByInvalidIDRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestDeleteMeeting_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestUpdateMeetingSettings_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestUpdateMeetingSettings_MeetingIDNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsMeetingIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestUpdateMeetingSettings_InvalidMeetingID(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsByInvalidMeetingIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestUpdateMeetingSettings_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestHandleParticipation_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response meetingsAPIMock.HandleParticipationResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.ParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertTrue(response.Data.HasNearMeeting, t)
}

func TestHandleParticipation_UserIDNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UserIDNotFoundParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestHandleParticipation_MeetingIDNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.MeetingIDNotFoundParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestHandleParticipation_InvalidIDs(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InvalidIDsParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestHandleParticipation_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.ParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestInviteUser_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestInviteUser_UserAlreadyInMeeting(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteAlreadyInMeetingUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserAlreadyInMeeting.Error(), response.ErrorDetail, t)
}

func TestInviteUser_MeetingIDNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserMeetingIDNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestInviteUser_InvalidIDs(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserMeetingInvalidIDsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestInviteUser_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}

func TestKickUser_Success(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.SuccessResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
	utils.AssertNil(response.Data, t)
}

func TestKickUser_UserNotInMeeting(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickNotInMeetingUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserNotInMeeting.Error(), response.ErrorDetail, t)
}

func TestKickUser_InvalidIDs(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickUserInvalidIDsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidID, response.ErrorDetail, t)
}

func TestKickUser_InternalError(t *testing.T) {
	mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.InternalError.Error(), response.ErrorDetail, t)
}
