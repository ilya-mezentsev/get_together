package repositories

import (
	"github.com/jmoiron/sqlx"
	"interfaces"
	"repositories/chat"
	"repositories/credentials"
	"repositories/decorators/logging"
	"repositories/meetings"
	"repositories/meetings_settings"
	"repositories/messages"
	"repositories/user_settings"
)

func Credentials(db *sqlx.DB) interfaces.CredentialsRepository {
	return logging.NewCredentialsRepositoryDecorator(credentials.New(db))
}

func Meetings(db *sqlx.DB) interfaces.FullMeetingsRepository {
	return logging.NewMeetingsRepositoryDecorator(meetings.New(db))
}

func MeetingsSettings(db *sqlx.DB) interfaces.MeetingsSettingsRepository {
	return logging.NewMeetingsSettingsRepositoryDecorator(meetings_settings.New(db))
}

func UserSettings(db *sqlx.DB) interfaces.UsersSettings {
	return logging.NewUserSettingsRepositoryDecorator(user_settings.New(db))
}

func Chat(db *sqlx.DB) interfaces.FullChatsRepository {
	return logging.NewChatRepositoryDecorator(chat.New(db))
}

func Messages(db *sqlx.DB) interfaces.Messages {
	return logging.NewMessagesRepositoryDecorator(messages.New(db))
}
