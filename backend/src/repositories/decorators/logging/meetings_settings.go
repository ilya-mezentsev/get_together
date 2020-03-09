package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type MeetingsSettingsRepositoryDecorator struct {
	repository interfaces.MeetingsSettingsRepository
}

func NewMeetingsSettingsRepositoryDecorator(
	repository interfaces.MeetingsSettingsRepository) MeetingsSettingsRepositoryDecorator {
	return MeetingsSettingsRepositoryDecorator{repository}
}

func (d MeetingsSettingsRepositoryDecorator) GetMeetingSettings(
	meetingId uint) (models.ParticipationMeetingSettings, error) {
	settings, err := d.repository.GetMeetingSettings(meetingId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting meeting settings: %v",
			Args: []interface{}{err},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
			},
		}, logger.Warning)
	}

	return settings, err
}

func (d MeetingsSettingsRepositoryDecorator) GetNearMeetings(
	data models.UserTimeCheckData) ([]models.TimeMeetingParameters, error) {
	parameters, err := d.repository.GetNearMeetings(data)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting near meeting: %v",
			Args: []interface{}{err},
			Optional: map[string]interface{}{
				"time_check_data": data,
			},
		}, logger.Warning)
	}

	return parameters, err
}
