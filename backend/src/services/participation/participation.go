package participation

import (
  "fmt"
  "interfaces"
  "internal_errors"
  "models"
  "plugins/meetings_time"
  "services"
)

const (
  bottomRatingValue = 60
  maxUsersCountReached = "max-users-count-reached"
  ageLessThanMin = "age-less-than-min"
  wrongGender = "wrong-gender"
)

type Service struct {
  userSettingsRepository interfaces.UsersSettingsRepository
  meetingsSettingsRepository interfaces.MeetingsSettingsRepository
}

func New(
  userSettingsRepository interfaces.UsersSettingsRepository,
  meetingsSettingsRepository interfaces.MeetingsSettingsRepository,
) Service {
  return Service{
    userSettingsRepository: userSettingsRepository,
    meetingsSettingsRepository: meetingsSettingsRepository,
  }
}

func (s Service) HandleParticipationRequest(
  request models.ParticipationRequest) (models.RejectInfo, interfaces.ErrorWrapper) {
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
  request models.ParticipationRequest) (models.FullUserInfo, models.ParticipationMeetingSettings, interfaces.ErrorWrapper) {
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
    return userSettings, meetingSettings, models.NewErrorWrapper(err, services.UserIdNotFound)
  default:
    return userSettings, meetingSettings, models.NewErrorWrapper(err, services.InternalError)
  }

  meetingSettings, err = s.meetingsSettingsRepository.GetMeetingSettings(request.MeetingId)
  switch err {
  case nil:
    break
  case internal_errors.UnableToFindByMeetingId:
    return userSettings, meetingSettings, models.NewErrorWrapper(err, services.MeetingIdNotFound)
  default:
    return userSettings, meetingSettings, models.NewErrorWrapper(err, services.InternalError)
  }

  return userSettings, meetingSettings, nil
}

func (s Service) hasNearMeeting(
  request models.ParticipationRequest,
  meetingSettings models.ParticipationMeetingSettings,
) (bool, interfaces.ErrorWrapper) {
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
    return false, models.NewErrorWrapper(err, services.UserIdNotFound)
  case internal_errors.UnableToFindByMeetingId:
    return false, models.NewErrorWrapper(err, services.MeetingIdNotFound)
  default:
    return false, models.NewErrorWrapper(err, services.InternalError)
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

  switch {
  case meetingSettings.UsersCount >= meetingSettings.MaxUsers:
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: maxUsersCountReached,
      Description: fmt.Sprintf("actual: %d", meetingSettings.UsersCount),
    })
  case userSettings.Age < meetingSettings.MinAge:
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: ageLessThanMin,
      Description: fmt.Sprintf("actual: %d, wanted: %d", userSettings.Age, meetingSettings.MinAge),
    })
  case userSettings.Gender != meetingSettings.Gender:
    inappropriateInfoFields = append(inappropriateInfoFields, models.InappropriateInfoField{
      ErrorCode: wrongGender,
      Description: fmt.Sprintf("actual: %s, wanted: %s", userSettings.Gender, meetingSettings.Gender),
    })
  }

  return inappropriateInfoFields
}
