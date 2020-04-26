package repositories

import (
	"models"
	"services/proxies/validation/plugins/validation"
	"time"
)

func GetFirstMessageSendingTime() time.Time {
	return GetAllMessages()[0].SendingTime
}

func GetMessageWithNotExistsChatId() models.Message {
	message := GetAllMessages()[0]
	message.ChatId = NotExistsChatId

	return message
}

func GetMessageWithNotExistsUserId() models.Message {
	message := GetAllMessages()[0]
	message.SenderId = GetNextUserId()

	return message
}

func GetAllMessages() []models.Message {
	var messages []models.Message
	for _, message := range ChatsMessages {
		sendingTime, _ := time.Parse(validation.DateFormat, message["sending_time"].(string))

		messages = append(messages, models.Message{
			ChatId:      uint(message["chat_id"].(int)),
			Text:        message["text"].(string),
			SenderId:    uint(message["sender_id"].(int)),
			SendingTime: sendingTime,
		})
	}

	return messages
}
