package meetings

import (
  mock "mock/services"
  "models"
  "services"
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
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
}

func TestService_GetExtendedMeetingsSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  meetings, err := service.GetExtendedMeetings(1)
  utils.AssertNil(err, t)
  expectedMeetings, _ := mock.MeetingsMockRepository.GetExtendedMeetings(models.UserMeetingStatusesData{
    UserId: 1,
    Invited: "",
    NotInvited: "",
  })
  utils.AssertEqual(expectedMeetings[0].PublicPlace.Latitude, meetings[0].PublicPlace.Latitude, t)
  utils.AssertEqual(expectedMeetings[0].PublicPlace.Longitude, meetings[0].PublicPlace.Longitude, t)
}

func TestService_GetExtendedMeetingsUserNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetExtendedMeetings(mock.NotExistsUserId)
  utils.AssertErrorsEqual(services.UserIdNotFound, err.ExternalError(), t)
}

func TestService_GetExtendedMeetingsInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetExtendedMeetings(mock.BadUserId)
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
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
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err.ExternalError(), t)
}

func TestService_GetFullMeetingInfoInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetFullMeetingInfo(mock.BadMeetingId)
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
}

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
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err.ExternalError(), t)
}

func TestService_DeleteMeetingInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.DeleteMeeting(mock.BadMeetingId)
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
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
  utils.AssertErrorsEqual(services.UserIdNotFound, err.ExternalError(), t)
}

func TestService_CreateMeetingInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.CreateMeeting(mock.BadUserId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
}

func TestService_UpdatedSettingsSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(1, mock.NewMeetingSettings)
  utils.AssertNil(err, t)
  meeting := mock.MeetingsMockRepository.Meetings[1]
  utils.AssertEqual(mock.NewMeetingSettings.PublicPlace, meeting.AllSettings.PublicPlace, t)
}

func TestService_UpdatedSettingsMeetingNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(mock.NotExistsMeetingId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err.ExternalError(), t)
}

func TestService_UpdatedSettings(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(mock.BadMeetingId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.InternalError, err.ExternalError(), t)
}
