package participation

import (
	mock "mock/services"
	"models"
	"services/errors"
	"testing"
	"utils"
)

var service = New(
	&mock.UsersSettingsRepository,
	&mock.MeetingsSettingsRepository,
)

func TestService_HandleParticipationRequestTooLowRatingTags(t *testing.T) {
	request, tags := mock.TooLowRatingTagsRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertTrue(mock.TagsEqual(tags, info.TooLowRatingTags), t)
}

func TestService_HandleParticipationRequestHasNearMeeting(t *testing.T) {
	info, err := service.HandleParticipationRequest(mock.HasNearMeetingRequest())

	utils.AssertNil(err, t)
	utils.AssertTrue(info.HasNearMeeting, t)
}

func TestService_HandleParticipationRequestInappropriateAge(t *testing.T) {
	request, inappropriateAgeField := mock.InappropriateAgeRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertTrue(mock.HasField(info.InappropriateInfoFields, inappropriateAgeField), t)
}

func TestService_HandleParticipationRequestWrongGender(t *testing.T) {
	request, inappropriateGenderField := mock.WrongGenderRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertTrue(mock.HasField(info.InappropriateInfoFields, inappropriateGenderField), t)
}

func TestService_HandleParticipationRequestMeetingWithoutGender(t *testing.T) {
	request, inappropriateGenderField := mock.MeetingWithoutGenderRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertFalse(mock.HasField(info.InappropriateInfoFields, inappropriateGenderField), t)
}

func TestService_HandleParticipationRequestMaxUsersCountReached(t *testing.T) {
	request, maxUsersCountReachedField := mock.MaxUsersCountReachedRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertTrue(mock.HasField(info.InappropriateInfoFields, maxUsersCountReachedField), t)
}

func TestService_HandleParticipationRequestNotExistsMeeting(t *testing.T) {
	_, err := service.HandleParticipationRequest(mock.NotExistsMeetingRequest())

	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_HandleParticipationRequestNotExistsUser(t *testing.T) {
	_, err := service.HandleParticipationRequest(mock.NotExistsUserRequest())

	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestService_HandleParticipationRequestInternalErrorBadMeetingId(t *testing.T) {
	_, err := service.HandleParticipationRequest(mock.InternalErrorBadMeetingIdRequest())

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_HasNearMeetingNotExistsMeeting(t *testing.T) {
	_, err := service.hasNearMeeting(mock.NotExistsMeetingRequest(), models.ParticipationMeetingSettings{})

	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_HasNearMeetingNotExistsUser(t *testing.T) {
	_, err := service.hasNearMeeting(mock.NotExistsUserRequest(), models.ParticipationMeetingSettings{})

	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestService_HandleParticipationRequestInternalErrorBadUserId(t *testing.T) {
	_, err := service.HandleParticipationRequest(mock.InternalErrorUserIdRequest())

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_HasNearMeetingInternalError(t *testing.T) {
	_, err := service.hasNearMeeting(mock.InternalErrorBadMeetingIdRequest(), models.ParticipationMeetingSettings{})

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_HandleParticipationRequestParticipationDescriptionRequiredTrue(t *testing.T) {
	request, descriptionRequiredField := mock.ParticipationWithoutDescriptionWhereItRequiredRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertTrue(mock.HasField(info.InappropriateInfoFields, descriptionRequiredField), t)
}

func TestService_HandleParticipationRequestParticipationDescriptionRequiredFalse(t *testing.T) {
	request, descriptionRequiredField := mock.ParticipationWithoutDescriptionWhereNotRequiredRequest()
	info, err := service.HandleParticipationRequest(request)

	utils.AssertNil(err, t)
	utils.AssertFalse(mock.HasField(info.InappropriateInfoFields, descriptionRequiredField), t)
}
