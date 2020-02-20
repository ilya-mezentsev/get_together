package meetings

import (
  "interfaces"
  "internal_errors"
  "models"
  "plugins/logger"
  "services"
  "services/meetings/plugins/coords"
)

const (
  invitedStatus = "invited"
  notInvitedStatus = "not-invited"
)

type Service struct {
  repository interfaces.MeetingsRepository
}

func New(repository interfaces.MeetingsRepository) Service {
  return Service{repository: repository}
}

func (s Service) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
  meeting, err := s.repository.GetFullMeetingInfo(meetingId)

  switch err {
  case nil:
    return meeting, nil
  case internal_errors.UnableToFindByMeetingId:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
      },
    }, logger.Warning)
    return models.PrivateMeeting{}, services.MeetingIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to get full meeting info: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
      },
    }, logger.Error)
    return models.PrivateMeeting{}, services.InternalError
  }
}

func (s Service) GetPublicMeetings() ([]models.PublicMeeting, error) {
  meetings, err := s.repository.GetPublicMeetings()

  switch err {
  case nil:
    return coords.ShakePublicMeetings(meetings), nil
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to get public meetings: %v",
      Args: []interface{}{err},
    }, logger.Error)
    return nil, services.InternalError
  }
}

func (s Service) GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error) {
  meetings, err := s.repository.GetExtendedMeetings(models.UserMeetingStatusesData{
    UserId: userId,
    Invited: invitedStatus,
    NotInvited: notInvitedStatus,
  })

  switch err {
  case nil:
    return coords.ShakeExtendedMeetings(meetings), nil
  case internal_errors.UnableToFindUserById:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Warning)
    return nil, services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to get extended meetings info: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "user_id": userId,
      },
    }, logger.Error)
    return nil, services.InternalError
  }
}

func (s Service) CreateMeeting(adminId uint, settings models.AllSettings) error {
  switch err := s.repository.CreateMeeting(adminId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "admin_id": adminId,
        "settings": settings,
      },
    }, logger.Warning)
    return services.UserIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to create meeting: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "admin_id": adminId,
        "settings": settings,
      },
    }, logger.Error)
    return services.InternalError
  }
}

func (s Service) DeleteMeeting(meetingId uint) error {
  switch err := s.repository.DeleteMeeting(meetingId); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
      },
    }, logger.Warning)
    return services.MeetingIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to delete meeting: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
      },
    }, logger.Error)
    return services.InternalError
  }
}

func (s Service) UpdatedSettings(meetingId uint, settings models.AllSettings) error {
  switch err := s.repository.UpdatedSettings(meetingId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    logger.WithFields(logger.Fields{
      MessageTemplate: err.Error(),
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
        "settings": settings,
      },
    }, logger.Warning)
    return services.MeetingIdNotFound
  default:
    logger.WithFields(logger.Fields{
      MessageTemplate: "unable to update meeting settings: %v",
      Args: []interface{}{err},
      Optional: map[string]interface{}{
        "meeting_id": meetingId,
        "settings": settings,
      },
    }, logger.Error)
    return services.InternalError
  }
}
