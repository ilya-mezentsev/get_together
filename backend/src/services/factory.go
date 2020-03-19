package services

import (
	"interfaces"
	"services/authentication"
	"services/meetings"
	"services/meetings_accessor"
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
