package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type MeetingsAccessorServiceProxy struct {
	service interfaces.MeetingsAccessorService
}

func NewMeetingsAccessorServiceProxy(service interfaces.MeetingsAccessorService) MeetingsAccessorServiceProxy {
	return MeetingsAccessorServiceProxy{service}
}

func (p MeetingsAccessorServiceProxy) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return models.PrivateMeeting{}, validationResults
	}

	return p.service.GetFullMeetingInfo(meetingId)
}

func (p MeetingsAccessorServiceProxy) GetPublicMeetings() ([]models.PublicMeeting, error) {
	return p.service.GetPublicMeetings()
}

func (p MeetingsAccessorServiceProxy) GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error) {
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return nil, validationResults
	}

	return p.service.GetExtendedMeetings(userId)
}
