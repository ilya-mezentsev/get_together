package messages

import (
	repositoriesMock "mock/repositories"
	mock "mock/services"
	"services/errors"
	"testing"
	"utils"
)

var service = New(&mock.MessagesMockRepository)

func TestService_SendSuccess(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	err := service.Save(repositoriesMock.GetAllMessages()[0])

	utils.AssertNil(err, t)
}

func TestService_SendChatNotFound(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	err := service.Save(repositoriesMock.GetMessageWithNotExistsChatId())

	utils.AssertErrorsEqual(errors.ChatIdNotFound, err, t)
}

func TestService_SendUserIdNotFound(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	err := service.Save(repositoriesMock.GetMessageWithNotExistsUserId())

	utils.AssertErrorsEqual(errors.UserIdNotFound, err, t)
}

func TestService_SendInternalError(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	err := service.Save(mock.GetMessageWithBadChatId())

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_GetLastMessagesSuccess(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	messagesCount := 2
	messages, err := service.GetLastMessages(1, uint(messagesCount))

	utils.AssertNil(err, t)
	utils.AssertTrue(len(messages) <= messagesCount && len(messages) > 0, t)
}

func TestService_GetLastMessagesInternalError(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	_, err := service.GetLastMessages(mock.BadChatId, 1)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}

func TestService_GetLastMessagesAfterSuccess(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	messagesCount := 2
	messages, err := service.GetLastMessagesAfter(1, 1, uint(messagesCount))

	utils.AssertNil(err, t)
	utils.AssertTrue(len(messages) <= messagesCount && len(messages) > 0, t)
}

func TestService_GetLastMessagesAfterInternalError(t *testing.T) {
	defer mock.MessagesMockRepository.ResetState()

	_, err := service.GetLastMessagesAfter(mock.BadChatId, 1, 1)

	utils.AssertErrorsEqual(errors.InternalError, err, t)
}
