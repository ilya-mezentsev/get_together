package repositories

import (
	"github.com/jmoiron/sqlx"
	"interfaces"
	"repositories/credentials"
	"repositories/decorators/logging"
	"repositories/meetings"
	"repositories/meetings_settings"
	"repositories/user_settings"
)

func Credentials(db *sqlx.DB) interfaces.CredentialsRepository {
	return logging.NewCredentialsRepositoryDecorator(credentials.New(db))
}

func Meetings(db *sqlx.DB) interfaces.Meetings {
	return logging.NewMeetingsRepositoryDecorator(meetings.New(db))
}

func MeetingsSettings(db *sqlx.DB) interfaces.MeetingsSettingsRepository {
	return logging.NewMeetingsSettingsRepositoryDecorator(meetings_settings.New(db))
}

func UserSettings(db *sqlx.DB) interfaces.UsersSettings {
	return logging.NewUserSettingsRepositoryDecorator(user_settings.New(db))
}
