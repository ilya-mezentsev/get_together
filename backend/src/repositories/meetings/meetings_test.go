package meetings

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"internal_errors"
	mock "mock/repositories"
	servicesMock "mock/services"
	"models"
	"os"
	"testing"
	"utils"
)

var (
	db         *sqlx.DB
	repository Repository
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

	mock.DropTables(db)
	repository = New(db)
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
	res := t.Run()
	mock.DropTables(db)
	os.Exit(res)
}

func TestRepository_GetFullMeetingInfoSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	info, err := repository.GetFullMeetingInfo(1)
	utils.AssertNil(err, t)
	utils.AssertEqual(mock.GetFirstLabeledPlace().Label, info.LabeledPlace.Label, t)
	utils.AssertEqual(mock.GetFirstLabeledPlace().GetLatitude(), info.LabeledPlace.GetLatitude(), t)
	utils.AssertEqual(mock.GetFirstLabeledPlace().GetLongitude(), info.LabeledPlace.GetLongitude(), t)
}

func TestRepository_GetFullMeetingInfoMeetingNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetFullMeetingInfo(mock.GetNotExistsMeetingId())
	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_GetFullMeetingInfoNoTableError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetFullMeetingInfo(mock.GetNotExistsMeetingId())
	utils.AssertNotNil(err, t)
}

func TestRepository_GetPublicMeetingsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	meetings, err := repository.GetPublicMeetings()
	utils.AssertNil(err, t)
	for idx, meeting := range meetings {
		utils.AssertEqual(mock.GetPlaceLongitudeById(idx), meeting.PublicPlace.Longitude, t)
		utils.AssertEqual(mock.GetPlaceLatitudeById(idx), meeting.PublicPlace.Latitude, t)
	}
}

func TestRepository_GetPublicMeetingsNoTableError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetPublicMeetings()
	utils.AssertNotNil(err, t)
}

func TestRepository_GetExtendedMeetingsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	meetings, err := repository.GetExtendedMeetings(models.UserMeetingStatusesData{
		UserId:     1,
		Invited:    "invited",
		NotInvited: "not-invited",
	})
	utils.AssertNil(err, t)
	for idx, meeting := range meetings {
		utils.AssertEqual(mock.FirstUserStatuses[idx], meeting.CurrentUserStatus, t)
	}
}

func TestRepository_GetExtendedMeetingsNoTableError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetExtendedMeetings(models.UserMeetingStatusesData{
		UserId:     1,
		Invited:    "invited",
		NotInvited: "not-invited",
	})
	utils.AssertNotNil(err, t)
}

func TestRepository_CreateMeetingSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateMeeting(1, servicesMock.NewMeetingSettings)
	utils.AssertNil(err, t)
	meeting, _ := repository.GetFullMeetingInfo(uint(len(mock.Meetings) + 1))
	utils.AssertEqual(servicesMock.NewMeetingSettings.Label, meeting.LabeledPlace.Label, t)
	utils.AssertEqual((&servicesMock.NewMeetingSettings).GetLatitude(), meeting.LabeledPlace.GetLatitude(), t)
	utils.AssertEqual((&servicesMock.NewMeetingSettings).GetLongitude(), meeting.LabeledPlace.GetLongitude(), t)
}

func TestRepository_CreateMeetingAdminNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateMeeting(mock.GetNotExistsUserId(), servicesMock.NewMeetingSettings)
	utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, t)
}

func TestRepository_CreateMeetingNoTableError(t *testing.T) {
	mock.DropTables(db)

	err := repository.CreateMeeting(mock.GetNotExistsMeetingId(), servicesMock.NewMeetingSettings)
	utils.AssertNotNil(err, t)
}

func TestRepository_DeleteMeetingSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.DeleteMeeting(1)
	utils.AssertNil(err, t)
	_, err = repository.GetFullMeetingInfo(1)
	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_DeleteMeetingMeetingNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.DeleteMeeting(mock.GetNotExistsMeetingId())
	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_DeleteMeetingNoTableError(t *testing.T) {
	mock.DropTables(db)

	err := repository.DeleteMeeting(mock.GetNotExistsMeetingId())
	utils.AssertNotNil(err, t)
}

func TestRepository_UpdatedSettingsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.UpdateSettings(2, servicesMock.NewMeetingSettings)
	utils.AssertNil(err, t)
	meeting, _ := repository.GetFullMeetingInfo(2)
	utils.AssertEqual(servicesMock.NewMeetingSettings.Title, meeting.Title, t)
}

func TestRepository_UpdatedSettingsMeetingNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.UpdateSettings(mock.GetNotExistsMeetingId(), servicesMock.NewMeetingSettings)
	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_UpdatedSettingsNoTableError(t *testing.T) {
	mock.DropTables(db)

	err := repository.UpdateSettings(mock.GetNotExistsMeetingId(), servicesMock.NewMeetingSettings)
	utils.AssertNotNil(err, t)
}

func TestRepository_AddUserToMeetingSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.AddUserToMeeting(1, mock.UserIdThatNotInFirstMeeting)
	userInMeeting, _ := repository.meetingHasUser(1, mock.UserIdThatNotInFirstMeeting)

	utils.AssertNil(err, t)
	utils.AssertTrue(userInMeeting, t)
}

func TestRepository_AddUserToMeetingUserAlreadyInMeetingError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.AddUserToMeeting(1, 1)
	utils.AssertErrorsEqual(internal_errors.UserAlreadyInMeeting, err, t)
}

func TestRepository_AddUserToMeetingNotExistsError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.AddUserToMeeting(0, 1)

	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_AddUserToMeetingInternalError(t *testing.T) {
	mock.DropTables(db)

	err := repository.AddUserToMeeting(1, 1)
	utils.AssertNotNil(err, t)
}

func TestRepository_KickUserFromMeetingSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.KickUserFromMeeting(1, 1)
	userInMeeting, _ := repository.meetingHasUser(1, 1)

	utils.AssertNil(err, t)
	utils.AssertFalse(userInMeeting, t)
}

func TestRepository_KickUserFromMeetingNotExistsError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.updateMeetingUserIds(KickUserFromMeetingQuery, 0, 1)

	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_KickUserFromMeetingUserNotInMeetingError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.KickUserFromMeeting(1, mock.UserIdThatNotInFirstMeeting)

	utils.AssertErrorsEqual(internal_errors.UserNotInMeeting, err, t)
}

func TestRepository_KickUserFromMeetingInternalError(t *testing.T) {
	mock.DropTables(db)

	err := repository.KickUserFromMeeting(1, 1)
	utils.AssertNotNil(err, t)
}

func TestRepository_UpdateMeetingUserIdsInternalError(t *testing.T) {
	mock.DropTables(db)

	err := repository.updateMeetingUserIds(AddUserIdToMeetingQuery, 1, 1)
	utils.AssertNotNil(err, t)
}
