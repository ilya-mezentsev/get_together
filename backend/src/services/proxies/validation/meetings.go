package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type MeetingsServiceProxy struct {
	service interfaces.Meetings
}

func NewMeetingsServiceProxy(service interfaces.Meetings) MeetingsServiceProxy {
	return MeetingsServiceProxy{service}
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

func (p MeetingsServiceProxy) UpdateSettings(meetingId uint, settings models.AllSettings) error {
	validationResults := p.validateAllSettings(settings)
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.UpdateSettings(meetingId, settings)
}

func (p MeetingsServiceProxy) AddUserToMeeting(meetingId, userId uint) error {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) ||
		!validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.AddUserToMeeting(meetingId, userId)
}

func (p MeetingsServiceProxy) KickUserFromMeeting(meetingId, userId uint) error {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) ||
		!validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.KickUserFromMeeting(meetingId, userId)
}
