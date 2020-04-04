package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type ChatRepositoryDecorator struct {
	repository interfaces.FullChatsRepository
}

func NewChatRepositoryDecorator(repository interfaces.FullChatsRepository) ChatRepositoryDecorator {
	return ChatRepositoryDecorator{repository}
}

func (d ChatRepositoryDecorator) GetMeetingChat(meetingId uint) (models.Chat, error) {
	chat, err := d.repository.GetMeetingChat(meetingId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting meeting chat by meeting id: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
			},
		}, logger.Warning)
	}

	return chat, err
}

func (d ChatRepositoryDecorator) GetUserChats(userId uint) ([]models.Chat, error) {
	chats, err := d.repository.GetUserChats(userId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting user chats by user id: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"user_id": userId,
			},
		}, logger.Warning)
	}

	return chats, err
}

func (d ChatRepositoryDecorator) CreateChat(meetingId uint, chatType string) error {
	err := d.repository.CreateChat(meetingId, chatType)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while creating chat: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"meeting_id": meetingId,
				"chat_type":  chatType,
			},
		}, logger.Warning)
	}

	return err
}

func (d ChatRepositoryDecorator) SetChatStatus(chatId uint, status string) error {
	err := d.repository.SetChatStatus(chatId, status)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while setting chat status: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"chat_id":     chatId,
				"chat_status": status,
			},
		}, logger.Warning)
	}

	return err
}
