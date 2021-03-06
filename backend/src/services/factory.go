package services

import (
	"interfaces"
	"services/authentication"
	"services/chat"
	"services/chat_accessor"
	"services/meetings"
	"services/meetings_accessor"
	"services/messages"
	"services/participation"
	"services/proxies/validation"
	"services/session"
	"services/user_settings"
)

func Authentication(repository interfaces.CredentialsRepository) interfaces.AuthenticationService {
	return validation.NewAuthenticationServiceProxy(authentication.New(repository))
}

func Meetings(repository interfaces.Meetings) interfaces.Meetings {
	return validation.NewMeetingsServiceProxy(meetings.New(repository))
}

func MeetingsAccessor(repository interfaces.MeetingsAccessorRepository) interfaces.MeetingsAccessorService {
	return validation.NewMeetingsAccessorServiceProxy(meetings_accessor.New(repository))
}

func Messages(repository interfaces.Messages) interfaces.Messages {
	return validation.NewMessagesProxy(messages.New(repository))
}

func Participation(
	userSettingsRepository interfaces.UsersSettings,
	meetingsSettingsRepository interfaces.MeetingsSettingsRepository,
) interfaces.ParticipationService {
	return validation.NewParticipationServiceProxy(
		participation.New(userSettingsRepository, meetingsSettingsRepository))
}

func Session(key string) interfaces.SessionService {
	return validation.NewSessionServiceProxy(session.New(key))
}

func UserSettings(repository interfaces.UsersSettings) interfaces.UsersSettings {
	return validation.NewUserSettingsServiceProxy(user_settings.New(repository))
}

func ChatAccessor(repository interfaces.ChatAccessor) interfaces.ChatAccessor {
	return validation.NewChatAccessorProxy(chat_accessor.New(repository))
}

func Chat(repository interfaces.ChatRepository) interfaces.Chat {
	return validation.NewChatProxy(chat.New(repository))
}
