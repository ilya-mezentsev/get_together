package participation

import (
  "io/ioutil"
  "log"
  mock "mock/services"
  "models"
  "os"
  "services"
  "testing"
  "utils"
)

var service = New(
  &mock.UsersSettingsRepository,
  &mock.MeetingsSettingsRepository,
)

func TestMain(m *testing.M) {
  log.SetOutput(ioutil.Discard)
  os.Exit(m.Run())
}

func TestService_HandleParticipationRequestBadRating(t *testing.T) {
  info, err := service.HandleParticipationRequest(mock.BadRatingRequest)

  utils.AssertNil(err, t)
  utils.AssertTrue(info.HasNearMeeting, t)
  utils.AssertTrue(mock.TagsEqual(mock.TagsWithBadRating, info.TooLowRatingTags), t)
  utils.AssertEqual(0, len(info.InappropriateInfoFields), t)
}

func TestService_HandleParticipationRequestMeetingFull(t *testing.T) {
  info, err := service.HandleParticipationRequest(mock.MeetingFullRequest)

  utils.AssertNil(err, t)
  utils.AssertEqual(mock.MaxUsersCountReached, info.InappropriateInfoFields[0], t)
}

func TestService_HandleParticipationRequestInappropriateAge(t *testing.T) {
  info, err := service.HandleParticipationRequest(mock.InappropriateAgeRequest)

  utils.AssertNil(err, t)
  utils.AssertEqual(mock.InappropriateAge, info.InappropriateInfoFields[0], t)
}

func TestService_HandleParticipationRequestWrongGender(t *testing.T) {
  info, err := service.HandleParticipationRequest(mock.WrongGenderRequest)

  utils.AssertNil(err, t)
  utils.AssertEqual(mock.WrongGender, info.InappropriateInfoFields[0], t)
}

func TestService_HandleParticipationRequestUserIdNotFound(t *testing.T) {
  _, err := service.HandleParticipationRequest(mock.NotExistsUserIdRequest)

  utils.AssertErrorsEqual(services.UserIdNotFound, err, t)
}

func TestService_HandleParticipationRequestMeetingIdNotFound(t *testing.T) {
  _, err := service.HandleParticipationRequest(mock.NotExistsMeetingIdRequest)

  utils.AssertErrorsEqual(services.MeetingIdNotFound, err, t)
}

func TestService_HandleParticipationRequestInternalError1(t *testing.T) {
  _, err := service.HandleParticipationRequest(mock.InternalErrorRequest1)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestService_HandleParticipationRequestInternalError2(t *testing.T) {
  _, err := service.HandleParticipationRequest(mock.InternalErrorRequest2)

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestService_HasNearMeetingRequestInternalError3(t *testing.T) {
  _, err := service.hasNearMeeting(mock.InternalErrorRequest2, models.ParticipationMeetingSettings{})

  utils.AssertErrorsEqual(services.InternalError, err, t)
}

func TestService_HasNearMeetingRequestMeetingNotFound(t *testing.T) {
  _, err := service.hasNearMeeting(mock.NotExistsMeetingIdRequest, models.ParticipationMeetingSettings{})

  utils.AssertErrorsEqual(services.MeetingIdNotFound, err, t)
}

func TestService_HasNearMeetingRequestUserNotFound(t *testing.T) {
  _, err := service.hasNearMeeting(mock.NotExistsUserIdRequest, models.ParticipationMeetingSettings{})

  utils.AssertErrorsEqual(services.UserIdNotFound, err, t)
}

func TestService_HandleParticipationRequestHasNearMeeting(t *testing.T) {
  info, err := service.HandleParticipationRequest(mock.HasNearMeetingRequest)

  utils.AssertNil(err, t)
  utils.AssertTrue(info.HasNearMeeting, t)
}
