package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type MeetingsServiceProxy struct {
	service interfaces.MeetingsService
}

func NewMeetingsServiceProxy(service interfaces.MeetingsService) MeetingsServiceProxy {
	return MeetingsServiceProxy{service}
}

func (p MeetingsServiceProxy) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return models.PrivateMeeting{}, validationResults
	}

	return p.service.GetFullMeetingInfo(meetingId)
}

func (p MeetingsServiceProxy) GetPublicMeetings() ([]models.PublicMeeting, error) {
	return p.service.GetPublicMeetings()
}

func (p MeetingsServiceProxy) GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error) {
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return nil, validationResults
	}

	return p.service.GetExtendedMeetings(userId)
}

func (p MeetingsServiceProxy) CreateMeeting(adminId uint, settings models.AllSettings) error {
	validationResults := p.validateAllSettings(settings)
	if !validation.ValidWholePositiveNumber(float64(adminId)) {
		validationResults.Add(InvalidID)
	}

	if validationResults.HasErrors() {
		return validationResults
	} else {
		return p.service.CreateMeeting(adminId, settings)
	}
}

func (p MeetingsServiceProxy) validateAllSettings(settings models.AllSettings) validationResults {
	validationResults := validationResults{}
	if !validation.ValidTitle(settings.Title) {
		validationResults.Add(InvalidMeetingTitle)
	}
	if !validation.ValidDescription(settings.Description) {
		validationResults.Add(InvalidMeetingDescription)
	}
	for _, tag := range settings.Tags {
		if !validation.ValidName(tag) {
			validationResults.Add(InvalidMeetingTag)
			break
		}
	}
	if !validation.ValidDate(settings.DateTime.Format(validation.DateFormat)) {
		validationResults.Add(InvalidMeetingDate)
	}
	if !validation.ValidWholePositiveNumber(float64(settings.MaxUsers)) {
		validationResults.Add(InvalidMeetingMaxUsers)
	}
	if !validation.ValidWholePositiveNumber(float64(settings.Duration)) {
		validationResults.Add(InvalidMeetingDuration)
	}
	if !validation.ValidWholePositiveNumber(float64(settings.MinAge)) {
		validationResults.Add(InvalidMeetingMinAge)
	}
	if !validation.ValidGender(settings.Gender) {
		validationResults.Add(InvalidMeetingGender)
	}

	return validationResults
}

func (p MeetingsServiceProxy) DeleteMeeting(meetingId uint) error {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.DeleteMeeting(meetingId)
}

func (p MeetingsServiceProxy) UpdatedSettings(meetingId uint, settings models.AllSettings) error {
	validationResults := p.validateAllSettings(settings)
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.UpdatedSettings(meetingId, settings)
}
