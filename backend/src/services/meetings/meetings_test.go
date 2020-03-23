package meetings

import (
	mock "mock/services"
	"services/errors"
	"testing"
	"utils"
)

var service = New(&mock.MeetingsMockRepository)

func TestService_DeleteMeetingSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.DeleteMeeting(1)
	utils.AssertNil(err, t)
	_, found := mock.MeetingsMockRepository.Meetings[0]
	utils.AssertTrue(!found, t)
}

func TestService_DeleteMeetingNotFoundError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.DeleteMeeting(mock.NotExistsMeetingId)
	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_DeleteMeetingInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.DeleteMeeting(mock.BadMeetingId)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_CreateMeetingSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.CreateMeeting(1, mock.NewMeetingSettings)
	utils.AssertNil(err, t)
	meeting := mock.MeetingsMockRepository.Meetings[2]
	utils.AssertEqual(mock.NewMeetingSettings.PublicPlace, meeting.AllSettings.PublicPlace, t)
}

func TestService_CreateMeetingUserIdNotFoundError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.CreateMeeting(mock.NotExistsUserId, mock.NewMeetingSettings)
	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestService_CreateMeetingInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.CreateMeeting(mock.BadUserId, mock.NewMeetingSettings)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_UpdateSettingsSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.UpdateSettings(1, mock.NewMeetingSettings)
	utils.AssertNil(err, t)
	meeting := mock.MeetingsMockRepository.Meetings[1]
	utils.AssertEqual(mock.NewMeetingSettings.PublicPlace, meeting.AllSettings.PublicPlace, t)
}

func TestService_UpdateSettingsMeetingNotFoundError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.UpdateSettings(mock.NotExistsMeetingId, mock.NewMeetingSettings)
	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_UpdateSettingsInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.UpdateSettings(mock.BadMeetingId, mock.NewMeetingSettings)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_AddUserToMeetingSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.AddUserToMeeting(1, mock.UserIdThatNotInFirstMeeting)
	utils.AssertNil(err, t)
	utils.AssertTrue(mock.HasUser(mock.MeetingsMockRepository.MeetingsUsers[1], mock.UserIdThatNotInFirstMeeting), t)
}

func TestService_AddUserToMeetingAlreadyInMeetingError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.AddUserToMeeting(1, 1)
	utils.AssertErrorsEqual(errors.UserAlreadyInMeeting, err, t)
}

func TestService_AddUserToMeetingNotFound(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.AddUserToMeeting(mock.NotExistsMeetingId, 1)
	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_AddUserToMeetingInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.AddUserToMeeting(mock.BadMeetingId, 1)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_KickUserFromMeetingSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.KickUserFromMeeting(1, 1)
	utils.AssertNil(err, t)
	utils.AssertFalse(mock.HasUser(mock.MeetingsMockRepository.MeetingsUsers[1], 1), t)
}

func TestService_KickUserFromMeetingUserNotInMeetingError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.KickUserFromMeeting(1, mock.UserIdThatNotInFirstMeeting)
	utils.AssertErrorsEqual(errors.UserNotInMeeting, err, t)
}

func TestService_KickUserFromMeetingInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	err := service.KickUserFromMeeting(mock.BadMeetingId, 1)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
