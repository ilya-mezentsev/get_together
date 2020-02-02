package meetings

import (
  "io/ioutil"
  "log"
  mock "mock/services"
  "os"
  "services"
  "testing"
  "utils"
)

var service = New(&mock.MeetingsMockRepository)

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestService_GetPublicMeetingsSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  meetings, err := service.GetPublicMeetings()
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  expectedMeetings, _ := mock.MeetingsMockRepository.GetPublicMeetings()
  utils.Assert(expectedMeetings[0].PublicPlace == meetings[0].PublicPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expectedMeetings, Got: meetings}))
    t.Fail()
  })
}

func TestService_GetPublicMeetingsInternalError(t *testing.T) {
  mock.MeetingsMockRepository.Meetings = nil
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetPublicMeetings()
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_GetExtendedMeetingsSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  meetings, err := service.GetExtendedMeetings(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  expectedMeetings, _ := mock.MeetingsMockRepository.GetExtendedMeetings(1)
  utils.Assert(expectedMeetings[0].PublicPlace == meetings[0].PublicPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expectedMeetings, Got: meetings}))
    t.Fail()
  })
}

func TestService_GetExtendedMeetingsUserNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetExtendedMeetings(mock.NotExistsUserId)
  utils.AssertErrorsEqual(services.UserIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_GetExtendedMeetingsInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetExtendedMeetings(mock.BadUserId)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_GetFullMeetingInfoSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  meeting, err := service.GetFullMeetingInfo(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  expectedMeeting, _ := mock.MeetingsMockRepository.GetFullMeetingInfo(1)
  utils.Assert(expectedMeeting.DefaultMeeting == meeting.DefaultMeeting, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expectedMeeting, Got: meeting}))
    t.Fail()
  })
}

func TestService_GetFullMeetingInfoMeetingNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetFullMeetingInfo(mock.NotExistsMeetingId)
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_GetFullMeetingInfoInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  _, err := service.GetFullMeetingInfo(mock.BadMeetingId)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_DeleteMeetingSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.DeleteMeeting(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  _, found := mock.MeetingsMockRepository.Meetings[0]
  utils.Assert(!found, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: false, Got: found}))
    t.Fail()
  })
}

func TestService_DeleteMeetingNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.DeleteMeeting(mock.NotExistsMeetingId)
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_DeleteMeetingInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.DeleteMeeting(mock.BadMeetingId)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_CreateMeetingSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.CreateMeeting(1, mock.NewMeetingSettings)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  meeting := mock.MeetingsMockRepository.Meetings[2]
  utils.Assert(mock.NewMeetingSettings.PublicPlace == meeting.AllSettings.PublicPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.NewMeetingSettings, Got: meeting}))
    t.Fail()
  })
}

func TestService_CreateMeetingUserIdNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.CreateMeeting(mock.NotExistsUserId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.UserIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_CreateMeetingInternalError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.CreateMeeting(mock.BadUserId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_UpdatedSettingsSuccess(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(1, mock.NewMeetingSettings)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  meeting := mock.MeetingsMockRepository.Meetings[1]
  utils.Assert(mock.NewMeetingSettings.PublicPlace == meeting.AllSettings.PublicPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.NewMeetingSettings, Got: meeting}))
    t.Fail()
  })
}

func TestService_UpdatedSettingsMeetingNotFoundError(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(mock.NotExistsMeetingId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.MeetingIdNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestService_UpdatedSettings(t *testing.T) {
  defer mock.MeetingsMockRepository.ResetState()

  err := service.UpdatedSettings(mock.BadMeetingId, mock.NewMeetingSettings)
  utils.AssertErrorsEqual(services.InternalError, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}
