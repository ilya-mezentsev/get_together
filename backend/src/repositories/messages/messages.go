package messages

import (
	"github.com/jmoiron/sqlx"
	"internal_errors"
	"models"
)

const (
	SaveMessageQuery     = `INSERT INTO messages(chat_id, sender_id, text) VALUES(:chat_id, :sender_id, :text)`
	GetLastMessagesQuery = `
	SELECT chat_id, sender_id, text, sending_time FROM messages
	WHERE chat_id = $1 ORDER BY sending_time DESC LIMIT $2`
	GetLastMessagesAfterQuery = `
	SELECT chat_id, sender_id, text, sending_time FROM messages
	WHERE chat_id = $1 AND sending_time >= (
		SELECT sending_time FROM messages WHERE id = $2
	) ORDER BY sending_time DESC LIMIT $3`

	chatIdNotFoundErrorMessage = `pq: insert or update on table "messages" violates foreign key constraint "messages_chat_id_fkey"`
	userIdNotFoundErrorMessage = `pq: insert or update on table "messages" violates foreign key constraint "messages_sender_id_fkey"`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Save(message models.Message) error {
	_, err := r.db.NamedExec(SaveMessageQuery, message)
	switch {
	case err == nil:
		return nil
	case err.Error() == chatIdNotFoundErrorMessage:
		return internal_errors.UnableToFindChatById
	case err.Error() == userIdNotFoundErrorMessage:
		return internal_errors.UnableToFindUserById
	default:
		return err
	}
}

func (r Repository) GetLastMessages(chatId, count uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Select(&messages, GetLastMessagesQuery, chatId, count)

	return messages, err
}

func (r Repository) GetLastMessagesAfter(chatId, messageId, count uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Select(&messages, GetLastMessagesAfterQuery, chatId, messageId, count)

	return messages, err
}
