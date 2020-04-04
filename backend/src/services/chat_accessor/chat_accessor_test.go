package chat_accessor

import (
	repositoriesMock "mock/repositories"
	mock "mock/services"
	"services/errors"
	"testing"
	"utils"
)

var service = New(&mock.ChatRepository)

func TestService_GetMeetingChatSuccess(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	chat, err := service.GetMeetingChat(1)

	utils.AssertNil(err, t)
	utils.AssertEqual(repositoriesMock.MeetingType, chat.Type, t)
}

func TestService_GetMeetingChatMeetingNotFound(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	_, err := service.GetMeetingChat(repositoriesMock.GetNotExistsMeetingId())

	utils.AssertErrorsEqual(errors.MeetingIdNotFound, err, t)
}

func TestService_GetMeetingChatInternalError(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	_, err := service.GetMeetingChat(mock.BadMeetingId)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_GetUserChatsSuccess(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	chats, err := service.GetUserChats(1)

	utils.AssertNil(err, t)
	utils.AssertNotNil(chats, t)
}

func TestService_GetUserChatsInternalError(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	_, err := service.GetUserChats(mock.BadUserId)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
