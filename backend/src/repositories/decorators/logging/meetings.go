package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type MeetingsRepositoryDecorator struct {
	repository interfaces.FullMeetingsRepository
}

func NewMeetingsRepositoryDecorator(repository interfaces.FullMeetingsRepository) MeetingsRepositoryDecorator {
	return MeetingsRepositoryDecorator{repository}
}

func (d MeetingsRepositoryDecorator) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
	meeting, err := d.repository.GetFullMeetingInfo(meetingId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting full meeting info: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
			},
		}, logger.Warning)
	}

	return meeting, err
}

func (d MeetingsRepositoryDecorator) GetPublicMeetings() ([]models.PublicMeeting, error) {
	meetings, err := d.repository.GetPublicMeetings()
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting public meetings: %v",
			Args: []interface{}{
				err,
			},
		}, logger.Warning)
	}

	return meetings, err
}

func (d MeetingsRepositoryDecorator) GetExtendedMeetings(
	userStatusesData models.UserMeetingStatusesData) ([]models.ExtendedMeeting, error) {
	meetings, err := d.repository.GetExtendedMeetings(userStatusesData)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting extended meetings: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"user_id": userStatusesData.UserId,
			},
		}, logger.Warning)
	}

	return meetings, err
}

func (d MeetingsRepositoryDecorator) CreateMeeting(adminId uint, settings models.AllSettings) error {
	err := d.repository.CreateMeeting(adminId, settings)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while creating meeting: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"admin_id": adminId,
				"settings": settings,
			},
		}, logger.Warning)
	}

	return err
}

func (d MeetingsRepositoryDecorator) DeleteMeeting(meetingId uint) error {
	err := d.repository.DeleteMeeting(meetingId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while deleting meeting: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
			},
		}, logger.Warning)
	}

	return err
}

func (d MeetingsRepositoryDecorator) UpdateSettings(meetingId uint, settings models.AllSettings) error {
	err := d.repository.UpdateSettings(meetingId, settings)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while updating settings: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
				"settings":   settings,
			},
		}, logger.Warning)
	}

	return err
}

func (d MeetingsRepositoryDecorator) AddUserToMeeting(meetingId, userId uint) error {
	err := d.repository.AddUserToMeeting(meetingId, userId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while adding user to meeting: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
				"user_id":    userId,
			},
		}, logger.Warning)
	}

	return err
}

func (d MeetingsRepositoryDecorator) KickUserFromMeeting(meetingId, userId uint) error {
	err := d.repository.KickUserFromMeeting(meetingId, userId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while kicking user from meeting: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
				"user_id":    userId,
			},
		}, logger.Warning)
	}

	return err
}
