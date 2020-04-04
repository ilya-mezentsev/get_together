package chat

import (
	"github.com/jmoiron/sqlx"
	"internal_errors"
	"models"
)

const (
	GetMeetingChatQuery = `
	SELECT id, type, status, created_at FROM chats
	WHERE meeting_id = $1 AND type = 'meeting' AND status != 'archived'`
	GetUserChatsQuery = `
	SELECT c.id, type, status, created_at FROM chats c
	JOIN messages m ON m.sender_id = $1 AND m.chat_id = c.id
	WHERE c.status != 'archived'`
	CreateChatQuery = `
	INSERT INTO chats(meeting_id, type)
	SELECT :meeting_id, :type
	WHERE NOT EXISTS (
		SELECT id FROM chats WHERE meeting_id = :meeting_id AND type = :type AND type = 'meeting'
	)`
	UpdateChatStatusQuery = `UPDATE chats SET status = :status WHERE id = :chat_id`

	meetingChatNotFound = `sql: no rows in result set`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) GetMeetingChat(meetingId uint) (models.Chat, error) {
	var chat models.Chat
	err := r.db.Get(&chat, GetMeetingChatQuery, meetingId)
	if err != nil && err.Error() == meetingChatNotFound {
		err = internal_errors.UnableToFindChatByMeetingId
	}

	return chat, err
}

func (r Repository) GetUserChats(userId uint) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.Select(&chats, GetUserChatsQuery, userId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (r Repository) CreateChat(meetingId uint, chatType string) error {
	res, err := r.db.NamedExec(CreateChatQuery, map[string]interface{}{
		"meeting_id": meetingId, "type": chatType,
	})
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return internal_errors.MeetingChatAlreadyExists
	}
	return nil
}

func (r Repository) SetChatStatus(chatId uint, status string) error {
	res, err := r.db.NamedExec(UpdateChatStatusQuery, map[string]interface{}{
		"chat_id": chatId, "status": status,
	})
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return internal_errors.UnableToFindChatById
	}
	return nil
}
