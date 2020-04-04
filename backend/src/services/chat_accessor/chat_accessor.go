package chat_accessor

import (
	"interfaces"
	"internal_errors"
	"models"
	"services/errors"
)

type Service struct {
	repository interfaces.ChatAccessor
}

func New(repository interfaces.ChatAccessor) Service {
	return Service{repository}
}

func (s Service) GetMeetingChat(meetingId uint) (models.Chat, error) {
	chat, err := s.repository.GetMeetingChat(meetingId)

	switch err {
	case nil:
		return chat, nil
	case internal_errors.UnableToFindChatByMeetingId:
		return models.Chat{}, errors.MeetingIdNotFound
	default:
		return models.Chat{}, errors.InternalError
	}
}

func (s Service) GetUserChats(userId uint) ([]models.Chat, error) {
	chats, err := s.repository.GetUserChats(userId)

	switch err {
	case nil:
		return chats, nil
	default:
		return nil, errors.InternalError
	}
}
