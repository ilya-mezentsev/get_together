package services

import (
	"fmt"
	"github.com/lib/pq"
	"internal_errors"
	"mock/repositories"
	"models"
	"services/proxies/validation/plugins/validation"
	"time"
)

type MeetingsSettingsRepositoryMock struct {
	meetings map[uint]models.ParticipationMeetingSettings
}

var (
	meetingIdToUsersCount      = map[uint]uint{1: 10, 2: 5, 3: 6}
	MeetingsSettingsRepository = MeetingsSettingsRepositoryMock{
		meetings: allMeetingsSettings(),
	}
)

func (m *MeetingsSettingsRepositoryMock) ResetState() {
	m.meetings = allMeetingsSettings()
}

func (m *MeetingsSettingsRepositoryMock) GetMeetingSettings(meetingId uint) (models.ParticipationMeetingSettings, error) {
	if meetingId == BadMeetingId {
		return models.ParticipationMeetingSettings{}, someInternalError
	} else if meetingId == repositories.GetNotExistsMeetingId() {
		return models.ParticipationMeetingSettings{}, internal_errors.UnableToFindMeetingById
	}

	return m.meetings[meetingId], nil
}

func (m *MeetingsSettingsRepositoryMock) GetNearMeetings(data models.UserTimeCheckData) ([]models.TimeMeetingParameters, error) {
	if data.MeetingId == BadMeetingId {
		return nil, someInternalError
	} else if data.MeetingId == repositories.GetNotExistsMeetingId() {
		return nil, internal_errors.UnableToFindMeetingById
	} else if data.UserId == repositories.GetNotExistsUserId() {
		return nil, internal_errors.UnableToFindUserById
	}

	var meetings []models.TimeMeetingParameters
	for meetingId, meeting := range allMeetingsSettings() {
		if meetingId != data.MeetingId {
			meetings = append(meetings, models.TimeMeetingParameters{
				DateTime: meeting.DateTime,
				Duration: meeting.Duration,
			})
		}
	}
	return meetings, nil
}

func allMeetingsSettings() map[uint]models.ParticipationMeetingSettings {
	settings := map[uint]models.ParticipationMeetingSettings{}
	for _, m := range repositories.MeetingsSettings {
		datetime, _ := time.Parse(validation.DateFormat, m["date_time"].(string))
		meetingId := uint(m["meeting_id"].(int))

		settings[meetingId] = models.ParticipationMeetingSettings{
			MeetingLimitations: models.MeetingLimitations{
				MaxUsers: uint(m["max_users"].(int)),
				Duration: uint(m["duration"].(int)),
				MinAge:   uint(m["min_age"].(int)),
				Gender:   m["gender"].(string),
			},
			MeetingParameters: models.MeetingParameters{
				DateTime:                   datetime,
				RequestDescriptionRequired: m["request_description_required"].(bool),
			},
			Tags:       pqStringArrayToStringArray(m["tags"].(*pq.StringArray)),
			UsersCount: meetingIdToUsersCount[meetingId],
		}
	}

	return settings
}

func pqStringArrayToStringArray(pqStringArray *pq.StringArray) []string {
	var strings []string
	for _, s := range *pqStringArray {
		strings = append(strings, s)
	}

	return strings
}

func TagsEqual(t1, t2 []string) bool {
	if len(t1) != len(t2) {
		return false
	}

	for idx := range t1 {
		if t1[idx] != t2[idx] {
			return false
		}
	}

	return true
}

func HasField(fields []models.InappropriateInfoField, field models.InappropriateInfoField) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}

	return false
}

func TooLowRatingTagsRequest() (models.ParticipationRequest, []string) {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: 1,
	}, []string{"tag2"}
}

func HasNearMeetingRequest() models.ParticipationRequest {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: 1,
	}
}

func InappropriateAgeRequest() (models.ParticipationRequest, models.InappropriateInfoField) {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: 1,
	}, inappropriateAge(12, 16)
}

func inappropriateAge(actual, wanted uint) models.InappropriateInfoField {
	return models.InappropriateInfoField{
		ErrorCode:   "age-less-than-min",
		Description: fmt.Sprintf("actual: %d, wanted: %d", actual, wanted),
	}
}

func WrongGenderRequest() (models.ParticipationRequest, models.InappropriateInfoField) {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: 2,
	}, wantedFemaleGender()
}

func wantedFemaleGender() models.InappropriateInfoField {
	return models.InappropriateInfoField{
		ErrorCode:   "wrong-gender",
		Description: "actual: male, wanted: female",
	}
}

func MaxUsersCountReachedRequest() (models.ParticipationRequest, models.InappropriateInfoField) {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: 3,
	}, maxUsersCountReached(6)
}

func maxUsersCountReached(actual uint) models.InappropriateInfoField {
	return models.InappropriateInfoField{
		ErrorCode:   "max-users-count-reached",
		Description: fmt.Sprintf("actual: %d", actual),
	}
}

func InternalErrorBadMeetingIdRequest() models.ParticipationRequest {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: BadMeetingId,
	}
}

func NotExistsMeetingRequest() models.ParticipationRequest {
	return models.ParticipationRequest{
		UserId:    1,
		MeetingId: repositories.GetNotExistsMeetingId(),
	}
}

func NotExistsUserRequest() models.ParticipationRequest {
	return models.ParticipationRequest{
		UserId:    repositories.GetNotExistsUserId(),
		MeetingId: 1,
	}
}

func InternalErrorUserIdRequest() models.ParticipationRequest {
	return models.ParticipationRequest{
		UserId:    BadUserId,
		MeetingId: 1,
	}
}

func ParticipationWithoutDescriptionWhereItRequiredRequest() (models.ParticipationRequest, models.InappropriateInfoField) {
	return models.ParticipationRequest{
			UserId:    1,
			MeetingId: 3,
		}, models.InappropriateInfoField{
			ErrorCode: "participation-request-description-required",
		}
}

func ParticipationWithoutDescriptionWhereNotRequiredRequest() (models.ParticipationRequest, models.InappropriateInfoField) {
	return models.ParticipationRequest{
			UserId:    1,
			MeetingId: 2,
		}, models.InappropriateInfoField{
			ErrorCode: "participation-request-description-required",
		}
}
