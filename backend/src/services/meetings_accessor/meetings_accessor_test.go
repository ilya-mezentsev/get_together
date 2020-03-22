package meetings_accessor

import (
	mock "mock/services"
	"models"
	"services/errors"
	"testing"
	"utils"
)

var service = New(&mock.MeetingsMockRepository)

func TestService_GetPublicMeetingsSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	meetings, err := service.GetPublicMeetings()
	utils.AssertNil(err, t)
	expectedMeetings, _ := mock.MeetingsMockRepository.GetPublicMeetings()
	utils.AssertEqual(expectedMeetings[0].PublicPlace.Latitude, meetings[0].PublicPlace.Latitude, t)
	utils.AssertEqual(expectedMeetings[0].PublicPlace.Longitude, meetings[0].PublicPlace.Longitude, t)
}

func TestService_GetPublicMeetingsInternalError(t *testing.T) {
	mock.MeetingsMockRepository.Meetings = nil
	defer mock.MeetingsMockRepository.ResetState()

	_, err := service.GetPublicMeetings()
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_GetExtendedMeetingsSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	meetings, err := service.GetExtendedMeetings(1)
	utils.AssertNil(err, t)
	expectedMeetings, _ := mock.MeetingsMockRepository.GetExtendedMeetings(models.UserMeetingStatusesData{
		UserId:     1,
		Invited:    "",
		NotInvited: "",
	})
	utils.AssertEqual(expectedMeetings[0].PublicPlace.Latitude, meetings[0].PublicPlace.Latitude, t)
	utils.AssertEqual(expectedMeetings[0].PublicPlace.Longitude, meetings[0].PublicPlace.Longitude, t)
}

func TestService_GetExtendedMeetingsUserNotFoundError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	_, err := service.GetExtendedMeetings(mock.NotExistsUserId)
	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestService_GetExtendedMeetingsInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	_, err := service.GetExtendedMeetings(mock.BadUserId)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_GetFullMeetingInfoSuccess(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	meeting, err := service.GetFullMeetingInfo(1)
	utils.AssertNil(err, t)
	expectedMeeting, _ := mock.MeetingsMockRepository.GetFullMeetingInfo(1)
	utils.AssertEqual(expectedMeeting.DefaultMeeting, meeting.DefaultMeeting, t)
}

func TestService_GetFullMeetingInfoMeetingNotFoundError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	_, err := service.GetFullMeetingInfo(mock.NotExistsMeetingId)
	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_GetFullMeetingInfoInternalError(t *testing.T) {
	defer mock.MeetingsMockRepository.ResetState()

	_, err := service.GetFullMeetingInfo(mock.BadMeetingId)
	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
