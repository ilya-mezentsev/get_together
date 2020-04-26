package services

import (
	"internal_errors"
	"mock/repositories"
	"models"
)

type MessagesRepositoryMock struct {
	chatId2Messages map[uint][]models.Message
}

var (
	MessagesMockRepository = MessagesRepositoryMock{
		chatId2Messages: getChatIdToMessages(),
	}
)

func (m *MessagesRepositoryMock) ResetState() {
	m.chatId2Messages = getChatIdToMessages()
}

func (m *MessagesRepositoryMock) Save(message models.Message) error {
	if message.ChatId == repositories.NotExistsChatId {
		return internal_errors.UnableToFindChatById
	} else if message.SenderId == repositories.GetNotExistsUserId() {
		return internal_errors.UnableToFindUserById
	} else if message.ChatId == BadChatId {
		return someInternalError
	}

	return nil
}

func (m *MessagesRepositoryMock) GetLastMessages(chatId, count uint) ([]models.Message, error) {
	if chatId == repositories.NotExistsChatId {
		return nil, internal_errors.UnableToFindChatById
	} else if chatId == BadChatId {
		return nil, someInternalError
	}

	messages := m.chatId2Messages[chatId]
	if int(count) < len(messages) {
		return messages[:count], nil
	} else {
		return messages, nil
	}
}

func (m *MessagesRepositoryMock) GetLastMessagesAfter(
	chatId, _, count uint) ([]models.Message, error) {
	return m.GetLastMessages(chatId, count)
}

func GetMessageWithBadChatId() models.Message {
	message := repositories.GetAllMessages()[0]
	message.ChatId = BadChatId

	return message
}

func getChatIdToMessages() map[uint][]models.Message {
	chatId2Messages := map[uint][]models.Message{}
	for _, message := range repositories.GetAllMessages() {
		messages, messagesFound := chatId2Messages[message.ChatId]
		if messagesFound {
			messages = append(messages, message)
		} else {
			chatId2Messages[message.ChatId] = []models.Message{message}
		}
	}

	return chatId2Messages
}
