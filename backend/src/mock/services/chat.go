package services

import (
	"internal_errors"
	"mock/repositories"
	"models"
)

type ChatRepositoryMock struct {
	meetingIdToChat map[uint]models.Chat
	userIdToChats   map[uint][]models.Chat
}

const BadChatId uint = 0

var ChatRepository = ChatRepositoryMock{
	meetingIdToChat: getMeetingIdToChat(),
	userIdToChats:   getUserIdToChats(),
}

func (m *ChatRepositoryMock) ResetState() {
	m.meetingIdToChat = getMeetingIdToChat()
	m.userIdToChats = getUserIdToChats()
}

func (m *ChatRepositoryMock) GetMeetingChat(meetingId uint) (models.Chat, error) {
	if meetingId == repositories.GetNotExistsMeetingId() {
		return models.Chat{}, internal_errors.UnableToFindChatByMeetingId
	} else if meetingId == BadMeetingId {
		return models.Chat{}, someInternalError
	}

	return m.meetingIdToChat[meetingId], nil
}

func (m *ChatRepositoryMock) GetUserChats(userId uint) ([]models.Chat, error) {
	if userId == BadUserId {
		return nil, someInternalError
	}

	return m.userIdToChats[userId], nil
}

func (m *ChatRepositoryMock) CreateChat(meetingId uint, chatType string) error {
	if meetingId == BadMeetingId {
		return someInternalError
	}

	chat, chatExists := m.meetingIdToChat[meetingId]
	if chatExists && chat.Type == repositories.MeetingType {
		return internal_errors.MeetingChatAlreadyExists
	}

	m.meetingIdToChat[meetingId] = models.Chat{Type: chatType}
	return nil
}

func (m *ChatRepositoryMock) SetChatStatus(chatId uint, status string) error {
	if chatId == repositories.NotExistsChatId {
		return internal_errors.UnableToFindChatById
	} else if chatId == BadChatId {
		return someInternalError
	}

	for meetingId, chat := range m.meetingIdToChat {
		if chat.Id == chatId {
			m.meetingIdToChat[meetingId] = models.Chat{
				Id:     chat.Id,
				Type:   chat.Type,
				Status: status,
			}
		}
	}

	return nil
}

func getMeetingIdToChat() map[uint]models.Chat {
	meetingIdToChat := map[uint]models.Chat{}
	for chatId, chat := range repositories.MeetingChats {
		meetingId, chatType := chat["meeting_id"].(int), chat["type"].(string)

		if chatType == repositories.MeetingType {
			meetingIdToChat[uint(meetingId)] = models.Chat{
				Id:   uint(chatId + 1),
				Type: chatType,
			}
		}
	}

	return meetingIdToChat
}

func getUserIdToChats() map[uint][]models.Chat {
	userIdToChats := map[uint][]models.Chat{}
	for _, message := range repositories.ChatsMessages {
		userId := uint(message["sender_id"].(int))
		_, userIdFound := userIdToChats[userId]

		if userIdFound {
			userIdToChats[userId] = append(userIdToChats[userId], models.Chat{
				Id: uint(message["chat_id"].(int)),
			})
		} else {
			userIdToChats[userId] = []models.Chat{
				{Id: uint(message["chat_id"].(int))},
			}
		}
	}

	return userIdToChats
}
