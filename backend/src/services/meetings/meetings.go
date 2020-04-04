package meetings

import (
	"interfaces"
	"internal_errors"
	"models"
	"services/errors"
)

type Service struct {
	repository interfaces.Meetings
}

func New(repository interfaces.Meetings) Service {
	return Service{repository}
}

func (s Service) CreateMeeting(adminId uint, settings models.AllSettings) error {
	switch s.repository.CreateMeeting(adminId, settings) {
	case nil:
		return nil
	case internal_errors.UnableToFindUserById:
		return errors.UserIdNotFound
	default:
		return errors.InternalError
	}
}

func (s Service) DeleteMeeting(meetingId uint) error {
	switch s.repository.DeleteMeeting(meetingId) {
	case nil:
		return nil
	case internal_errors.UnableToFindMeetingById:
		return errors.MeetingIdNotFound
	default:
		return errors.InternalError
	}
}

func (s Service) UpdateSettings(meetingId uint, settings models.AllSettings) error {
	switch s.repository.UpdateSettings(meetingId, settings) {
	case nil:
		return nil
	case internal_errors.UnableToFindMeetingById:
		return errors.MeetingIdNotFound
	default:
		return errors.InternalError
	}
}

func (s Service) AddUserToMeeting(meetingId, userId uint) error {
	switch s.repository.AddUserToMeeting(meetingId, userId) {
	case nil:
		return nil
	case internal_errors.UserAlreadyInMeeting:
		return errors.UserAlreadyInMeeting
	case internal_errors.UnableToFindMeetingById:
		return errors.MeetingIdNotFound
	default:
		return errors.InternalError
	}
}

func (s Service) KickUserFromMeeting(meetingId, userId uint) error {
	switch s.repository.KickUserFromMeeting(meetingId, userId) {
	case nil:
		return nil
	case internal_errors.UserNotInMeeting:
		return errors.UserNotInMeeting
	default:
		return errors.InternalError
	}
}
