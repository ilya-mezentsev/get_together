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
)

type Service struct {
  repository interfaces.MeetingsRepository
}

func New(repository interfaces.MeetingsRepository) Service {
  return Service{repository}
}

func (s Service) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
  meeting, err := s.repository.GetFullMeetingInfo(meetingId)

  switch err {
  case nil:
    return meeting, nil
  case internal_errors.UnableToFindByMeetingId:
    return models.PrivateMeeting{}, services.MeetingIdNotFound
  default:
    return models.PrivateMeeting{}, services.InternalError
  }
}

func (s Service) GetPublicMeetings() ([]models.PublicMeeting, error) {
  meetings, err := s.repository.GetPublicMeetings()

  switch err {
  case nil:
    return coords.ShakePublicMeetings(meetings), nil
  default:
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
    return nil, services.UserIdNotFound
  default:
    return nil, services.InternalError
  }
}

func (s Service) CreateMeeting(adminId uint, settings models.AllSettings) error {
  switch err := s.repository.CreateMeeting(adminId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindUserById:
    return services.UserIdNotFound
  default:
    return services.InternalError
  }
}

func (s Service) DeleteMeeting(meetingId uint) error {
  switch err := s.repository.DeleteMeeting(meetingId); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    return services.MeetingIdNotFound
  default:
    return services.InternalError
  }
}

func (s Service) UpdatedSettings(meetingId uint, settings models.AllSettings) error {
  switch err := s.repository.UpdatedSettings(meetingId, settings); err {
  case nil:
    return nil
  case internal_errors.UnableToFindByMeetingId:
    return services.MeetingIdNotFound
  default:
    return services.InternalError
  }
}
