package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type ChatAccessorProxy struct {
	service interfaces.ChatAccessor
}

func NewChatAccessorProxy(service interfaces.ChatAccessor) ChatAccessorProxy {
	return ChatAccessorProxy{service}
}

func (p ChatAccessorProxy) GetMeetingChat(meetingId uint) (models.Chat, error) {
	if !validation.ValidWholePositiveNumber(float64(meetingId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return models.Chat{}, validationResults
	}

	return p.service.GetMeetingChat(meetingId)
}

func (p ChatAccessorProxy) GetUserChats(userId uint) ([]models.Chat, error) {
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidId)
		return nil, validationResults
	}

	return p.service.GetUserChats(userId)
}
