package meetings

import (
  "interfaces"
  "internal_errors"
  "models"
  "services"
  "services/meetings/plugins/coords"
)

const (
  invitedStatus = "invited"
  notInvitedStatus = "not-invited"
  defaultDuration  = 4
)

type Service struct {
  repository interfaces.MeetingsRepository
}

func New(repository interfaces.MeetingsRepository) Service {
  return Service{repository: repository}
}

func (s Service) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, interfaces.ErrorWrapper) {
  meeting, err := s.repository.GetFullMeetingInfo(meetingId)

  switch err {
  case nil:
    return meeting, nil
  case internal_errors.UnableToFindByMeetingId:
    return models.PrivateMeeting{}, models.NewErrorWrapper(err, services.MeetingIdNotFound)
  default:
    return models.PrivateMeeting{}, models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) GetPublicMeetings() ([]models.PublicMeeting, interfaces.ErrorWrapper) {
  meetings, err := s.repository.GetPublicMeetings()

  switch err {
  case nil:
    return coords.ShakePublicMeetings(meetings), nil
  default:
    return nil, models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, interfaces.ErrorWrapper) {
  meetings, err := s.repository.GetExtendedMeetings(models.UserMeetingStatusesData{
    UserId: userId,
    Invited: invitedStatus,
    NotInvited: notInvitedStatus,
  })

  switch err {
  case nil:
    return coords.ShakeExtendedMeetings(meetings), nil
  case internal_errors.UnableToFindUserById:
    return nil, models.NewErrorWrapper(err, services.UserIdNotFound)
  default:
    return nil, models.NewErrorWrapper(err, services.InternalError)
  }
}

func changeMeetingDurationIfNeeded(m *models.MeetingLimitation) {
  if m.Duration == 0 {
    m.Duration = defaultDuration
  }
}

func (s Service) CreateMeeting(adminId uint, settings models.AllSettings) interfaces.ErrorWrapper {
  switch changeMeetingDurationIfNeeded(settings.MeetingLimitation);  /*тип так? -_- */
  err := s.repository.CreateMeeting(adminId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    return models.NewErrorWrapper(err, services.UserIdNotFound)
  default:
    return models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) DeleteMeeting(meetingId uint) interfaces.ErrorWrapper {
  switch err := s.repository.DeleteMeeting(meetingId); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    return models.NewErrorWrapper(err, services.MeetingIdNotFound)
  default:
    return models.NewErrorWrapper(err, services.InternalError)
  }
}

func (s Service) UpdatedSettings(meetingId uint, settings models.AllSettings) interfaces.ErrorWrapper {
  switch err := s.repository.UpdatedSettings(meetingId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    return models.NewErrorWrapper(err, services.MeetingIdNotFound)
  default:
    return models.NewErrorWrapper(err, services.InternalError)
  }
}
