package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type MessagesRepositoryDecorator struct {
	repository interfaces.Messages
}

func NewMessagesRepositoryDecorator(repository interfaces.Messages) MessagesRepositoryDecorator {
	return MessagesRepositoryDecorator{repository}
}

func (d MessagesRepositoryDecorator) Save(message models.Message) error {
	err := d.repository.Save(message)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while saving message: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"message": message,
			},
		}, logger.Warning)
	}

	return err
}

func (d MessagesRepositoryDecorator) GetLastMessages(chatId, count uint) ([]models.Message, error) {
	messages, err := d.repository.GetLastMessages(chatId, count)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting messages: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"chat_id": chatId,
				"count":   count,
			},
		}, logger.Warning)
	}

	return messages, err
}

func (d MessagesRepositoryDecorator) GetLastMessagesAfter(
	chatId, messageId, count uint,
) ([]models.Message, error) {
	messages, err := d.repository.GetLastMessagesAfter(chatId, messageId, count)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting messages after date: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"chat_id":    chatId,
				"message_id": messageId,
				"count":      count,
			},
		}, logger.Warning)
	}

	return messages, err
}
