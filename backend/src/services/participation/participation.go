package participation

import (
  "fmt"
  "interfaces"
  "internal_errors"
  "models"
  "plugins/logger"
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
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Warning)
    return userSettings, meetingSettings, services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "error while getting user settings: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Error)
    return userSettings, meetingSettings, services.InternalError
  }

  meetingSettings, err = s.meetingsSettingsRepository.GetMeetingSettings(request.MeetingId)
  switch err {
  case nil:
    break
  case internal_errors.UnableToFindByMeetingId:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Warning)
    return userSettings, meetingSettings, services.MeetingIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "error while getting meeting settings: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Error)
    return userSettings, meetingSettings, services.InternalError
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
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Warning)
    return false, services.UserIdNotFound
  case internal_errors.UnableToFindByMeetingId:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Warning)
    return false, services.MeetingIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "error while finding near meetings: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "request": request,
      },
    }, logger.Warning)
    return false, services.InternalError
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