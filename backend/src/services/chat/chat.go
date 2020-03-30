package chat

import (
	"interfaces"
	"internal_errors"
	"services/errors"
)

const (
	meetingChatType        = "meeting"
	meetingRequestChatType = "meeting_request"
	archivedChatStatus     = "archived"
)

type Service struct {
	repository interfaces.ChatRepository
}

func New(repository interfaces.ChatRepository) Service {
	return Service{repository}
}

func (s Service) CreateMeetingChat(meetingId uint) error {
	switch err := s.repository.CreateChat(meetingId, meetingChatType); err {
	case nil:
		return nil
	default:
		return errors.InternalError
	}
}

func (s Service) CreateMeetingRequestChat(meetingId uint) error {
	switch err := s.repository.CreateChat(meetingId, meetingRequestChatType); err {
	case nil:
		return nil
	default:
		return errors.InternalError
	}
}

func (s Service) CloseChat(chatId uint) error {
	switch err := s.repository.SetChatStatus(chatId, archivedChatStatus); err {
	case nil:
		return nil
	case internal_errors.UnableToFindChatById:
		return errors.ChatIdNotFound
	default:
		return errors.InternalError
	}
}
