package participation

import (
	"fmt"
	"interfaces"
	"internal_errors"
	"models"
	"services/errors"
	"services/participation/plugins/meetings_time"
)

const (
  bottomRatingValue = 60
  maxUsersCountReached = "max-users-count-reached"
  ageLessThanMin = "age-less-than-min"
  wrongGender = "wrong-gender"
)

type Service struct {
  userSettingsRepository interfaces.UsersSettings
  meetingsSettingsRepository interfaces.MeetingsSettingsRepository
}

func New(
  userSettingsRepository interfaces.UsersSettings,
  meetingsSettingsRepository interfaces.MeetingsSettingsRepository,
) Service {
  return Service{userSettingsRepository, meetingsSettingsRepository}
}

func (s Service) HandleParticipationRequest(request models.ParticipationRequest) (models.RejectInfo, error) {
  userSettings, meetingSettings, err := s.getUserAndMeetingSettings(request)
  if err != nil {
    return models.RejectInfo{}, err
  }

  hasNearMeeting, err := s.hasNearMeeting(request, meetingSettings)
  if err != nil {
    return models.RejectInfo{}, err
  }

  return models.RejectInfo{
    TooLowRatingTags: s.getTooLowRatingTags(userSettings, meetingSettings),
    InappropriateInfoFields: s.parseUserAndMeetingSettings(userSettings, meetingSettings),
    HasNearMeeting: hasNearMeeting,
  }, nil
}

func (s Service) getUserAndMeetingSettings(
  request models.ParticipationRequest) (models.FullUserInfo, models.ParticipationMeetingSettings, error) {
  var (
    userSettings models.FullUserInfo
    meetingSettings models.ParticipationMeetingSettings
    err error
  )

  userSettings, err = s.userSettingsRepository.GetUserSettings(request.UserId)
  switch err {
  case nil:
    break
  case internal_errors.UnableToFindUserById:
    return userSettings, meetingSettings, errors.UserIdNotFound
  default:
    return userSettings, meetingSettings, errors.InternalError
  }

  meetingSettings, err = s.meetingsSettingsRepository.GetMeetingSettings(request.MeetingId)
  switch err {
  case nil:
    break
  case internal_errors.UnableToFindByMeetingId:
    return userSettings, meetingSettings, errors.MeetingIdNotFound
  default:
    return userSettings, meetingSettings, errors.InternalError
  }

  return userSettings, meetingSettings, nil
}

func (s Service) hasNearMeeting(
  request models.ParticipationRequest,
  meetingSettings models.ParticipationMeetingSettings,
) (bool, error) {
  meetings, err := s.meetingsSettingsRepository.GetNearMeetings(
    models.UserTimeCheckData{
      UserId: request.UserId,
      MeetingId: request.MeetingId,
    })

  switch err {
  case nil:
    return meetings_time.MeetingsNearTo(
      models.TimeMeetingParameters{
        DateTime: meetingSettings.DateTime,
        Duration: meetingSettings.Duration,
      }, meetings), nil
  case internal_errors.UnableToFindUserById:
    return false, errors.UserIdNotFound
  case internal_errors.UnableToFindByMeetingId:
    return false, errors.MeetingIdNotFound
  default:
    return false, errors.InternalError
  }
}

func (s Service) getTooLowRatingTags(
  userSettings models.FullUserInfo,
  meetingSettings models.ParticipationMeetingSettings,
) []string {
  var tooLowRatingTags []string
  meetingTags := s.getExistingTagsMap(meetingSettings.Tags)
  for _, rating := range userSettings.Rating {
    _, foundTag := meetingTags[rating.Tag]
    if !foundTag {
      continue
    }

    if rating.Value < bottomRatingValue {
      tooLowRatingTags = append(tooLowRatingTags, rating.Tag)
    }
  }

  return tooLowRatingTags
}

func (s Service) getExistingTagsMap(tags []string) map[string]bool {
  mapTags := make(map[string]bool)
  for _, tag := range tags {
    mapTags[tag] = true
  }

  return mapTags
}

func (s Service) parseUserAndMeetingSettings(
  userSettings models.FullUserInfo,
  meetingSettings models.ParticipationMeetingSettings,
) []models.InappropriateInfoField {
  var inappropriateInfoFields []models.InappropriateInfoField

  if meetingSettings.UsersCount >= meetingSettings.MaxUsers {
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: maxUsersCountReached,
      Description: fmt.Sprintf("actual: %d", meetingSettings.UsersCount),
    })
  }

  if userSettings.Age < meetingSettings.MinAge {
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: ageLessThanMin,
      Description: fmt.Sprintf("actual: %d, wanted: %d", userSettings.Age, meetingSettings.MinAge),
    })
  }

  if userSettings.Gender != meetingSettings.Gender {
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: wrongGender,
      Description: fmt.Sprintf("actual: %s, wanted: %s", userSettings.Gender, meetingSettings.Gender),
    })
  }

  return inappropriateInfoFields
}
