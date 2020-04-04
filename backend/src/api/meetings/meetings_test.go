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

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestCreateMeeting_NotExistsAdminId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingNotExistsAdminIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestCreateMeeting_InvalidAdminId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.CreateMeetingInvalidAdminIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestDeleteMeeting_MeetingIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestDeleteMeeting_InvalidMeetingIdError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.DeleteMeetingByInvalidIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
}

func TestUpdateMeetingSettings_MeetingIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsMeetingIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestUpdateMeetingSettings_InvalidMeetingId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UpdateMeetingSettingsByInvalidMeetingIdRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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

func TestHandleParticipation_UserIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.UserIdNotFoundParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.UserIdNotFound.Error(), response.ErrorDetail, t)
}

func TestHandleParticipation_MeetingIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.MeetingIdNotFoundParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestHandleParticipation_InvalidIds(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InvalidIdsParticipationRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
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

func TestInviteUser_MeetingIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserMeetingIdNotFoundRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(errors.MeetingIdNotFound.Error(), response.ErrorDetail, t)
}

func TestInviteUser_InvalidIds(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.InviteUserMeetingInvalidIdsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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

	var response models.DefaultResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickUserRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusOk, response.Status, t)
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

func TestKickUser_InvalidIds(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	var response models.ErrorResponse
	err := json.NewDecoder(
		utils.MakeRequest(meetingsAPIMock.KickUserInvalidIdsRequest(router))).Decode(&response)

	utils.AssertNil(err, t)
	utils.AssertEqual(api.StatusError, response.Status, t)
	utils.AssertEqual(validation.InvalidId, response.ErrorDetail, t)
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
