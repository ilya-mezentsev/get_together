package validation

import (
	"interfaces"
	"services/proxies/validation/plugins/validation"
)

type ChatProxy struct {
	service interfaces.Chat
}

func NewChatProxy(service interfaces.Chat) ChatProxy {
	return ChatProxy{service}
}

func (p ChatProxy) CreateMeetingChat(meetingId uint) error {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return validationResults
	}

	return p.service.CreateMeetingChat(meetingId)
}

func (p ChatProxy) CreateMeetingRequestChat(meetingId uint) error {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return validationResults
	}

	return p.service.CreateMeetingRequestChat(meetingId)
}

func (p ChatProxy) CloseChat(chatId uint) error {
	if !validation.ValidWholePositiveNumber(float64(chatId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return validationResults
	}

	return p.service.CloseChat(chatId)
}
