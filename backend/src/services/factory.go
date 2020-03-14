package services

import (
	"interfaces"
	"services/authentication"
	"services/meetings"
	"services/participation"
	"services/proxies/validation"
	"services/session"
	"services/user_settings"
)

func Authentication(repository interfaces.CredentialsRepository) interfaces.AuthenticationService {
	return validation.NewAuthenticationServiceProxy(authentication.New(repository))
}

func Meetings(repository interfaces.MeetingsRepository) interfaces.MeetingsService {
	return validation.NewMeetingsServiceProxy(meetings.New(repository))
}

func Participation(
	userSettingsRepository interfaces.UsersSettingsRepository,
	meetingsSettingsRepository interfaces.MeetingsSettingsRepository,
) interfaces.ParticipationService {
	return validation.NewParticipationServiceProxy(
		participation.New(userSettingsRepository, meetingsSettingsRepository))
}

func Session(key string) interfaces.SessionService {
	return validation.NewSessionServiceProxy(session.New(key))
}

func UserSettings(repository interfaces.UsersSettingsRepository) interfaces.UserSettingsService {
	return validation.NewUserSettingsServiceProxy(user_settings.New(repository))
}
