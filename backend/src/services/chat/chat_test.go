package chat

import (
	repositoriesMock "mock/repositories"
	mock "mock/services"
	"services/errors"
	"testing"
	"utils"
)

var service = New(&mock.ChatRepository)

func TestService_CreateMeetingChatSuccess(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CreateMeetingChat(repositoriesMock.MeetingIdWithoutMeetingChat)
	chat, _ := mock.ChatRepository.GetMeetingChat(repositoriesMock.MeetingIdWithoutMeetingChat)

	utils.AssertNil(err, t)
	utils.AssertEqual(meetingChatType, chat.Type, t)
}

func TestService_CreateMeetingChatInternalError(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CreateMeetingChat(mock.BadMeetingId)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_CreateMeetingRequestChatSuccess(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CreateMeetingRequestChat(repositoriesMock.MeetingIdWithoutMeetingChat)
	chat, _ := mock.ChatRepository.GetMeetingChat(repositoriesMock.MeetingIdWithoutMeetingChat)

	utils.AssertNil(err, t)
	utils.AssertEqual(meetingRequestChatType, chat.Type, t)
}

func TestService_CreateMeetingRequestChatInternalError(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CreateMeetingRequestChat(mock.BadMeetingId)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_CloseChatSuccess(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CloseChat(1)
	chat, _ := mock.ChatRepository.GetMeetingChat(1)

	utils.AssertNil(err, t)
	utils.AssertEqual(archivedChatStatus, chat.Status, t)
}

func TestService_CloseChatNotFound(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CloseChat(repositoriesMock.NotExistsChatId)

	utils.AssertErrorsEqual(errors.ChatIdNotFound, err, t)
}

func TestService_CloseChatInternalError(t *testing.T) {
	defer mock.ChatRepository.ResetState()

	err := service.CloseChat(mock.BadChatId)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
