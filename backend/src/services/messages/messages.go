package messages

import (
	"interfaces"
	"internal_errors"
	"models"
	"services/errors"
)

type Service struct {
	repository interfaces.Messages
}

func New(repository interfaces.Messages) Service {
	return Service{repository}
}

func (s Service) Save(message models.Message) error {
	err := s.repository.Save(message)

	switch {
	case err == nil:
		return nil
	case err == internal_errors.UnableToFindChatById:
		return errors.ChatIdNotFound
	case err == internal_errors.UnableToFindUserById:
		return errors.UserIdNotFound
	default:
		return errors.InternalError
	}
}

func (s Service) GetLastMessages(chatId, count uint) ([]models.Message, error) {
	messages, err := s.repository.GetLastMessages(chatId, count)

	switch err {
	case nil:
		return messages, nil
	default:
		return nil, errors.InternalError
	}
}

func (s Service) GetLastMessagesAfter(chatId, messageId, count uint) ([]models.Message, error) {
	messages, err := s.repository.GetLastMessagesAfter(chatId, messageId, count)

	switch err {
	case nil:
		return messages, nil
	default:
		return nil, errors.InternalError
	}
}
