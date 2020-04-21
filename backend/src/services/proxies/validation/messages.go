package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type MessagesProxy struct {
	service interfaces.Messages
}

func NewMessagesProxy(service interfaces.Messages) MessagesProxy {
	return MessagesProxy{service}
}

func (p MessagesProxy) Save(message models.Message) error {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(message.ChatId)) ||
		!validation.ValidWholePositiveNumber(float64(message.SenderId)) {
		validationResults.Add(InvalidId)
	}
	if !validation.ValidDate(message.SendingTime.Format(validation.DateFormat)) {
		validationResults.Add(InvalidDate)
	}
	if !validation.ValidMessage(message.Text) {
		validationResults.Add(InvalidMessageText)
	}

	if validationResults.HasErrors() {
		return validationResults
	} else {
		return p.service.Save(message)
	}
}

func (p MessagesProxy) GetLastMessages(chatId, count uint) ([]models.Message, error) {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(chatId)) {
		validationResults.Add(InvalidId)
	}
	if !validation.ValidWholePositiveNumber(float64(count)) {
		validationResults.Add(InvalidCount)
	}

	if validationResults.HasErrors() {
		return nil, validationResults
	} else {
		return p.service.GetLastMessages(chatId, count)
	}
}

func (p MessagesProxy) GetLastMessagesAfter(chatId, messageId, count uint) ([]models.Message, error) {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(chatId)) ||
		!validation.ValidWholePositiveNumber(float64(messageId)) {
		validationResults.Add(InvalidId)
	}
	if !validation.ValidWholePositiveNumber(float64(count)) {
		validationResults.Add(InvalidCount)
	}

	if validationResults.HasErrors() {
		return nil, validationResults
	} else {
		return p.service.GetLastMessagesAfter(chatId, messageId, count)
	}
}
