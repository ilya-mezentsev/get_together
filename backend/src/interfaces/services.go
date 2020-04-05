package interfaces

import (
	"models"
	"net/http"
)

type (
	AuthenticationService interface {
		RegisterUser(credentials models.UserCredentials) error
		Login(credentials models.UserCredentials) (models.UserSession, error)
		ChangePassword(userId uint, password string) error
	}

	SessionAccessorService interface {
		GetSession(r *http.Request) (map[string]interface{}, error)
	}

	SessionService interface {
		SessionAccessorService
		SetSession(r *http.Request, session models.UserSession) error
		InvalidateSession(r *http.Request)
	}

	MeetingsAccessorService interface {
		GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error)
		GetPublicMeetings() ([]models.PublicMeeting, error)
		GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error)
	}

	Meetings interface {
		CreateMeeting(adminId uint, settings models.AllSettings) error
		DeleteMeeting(meetingId uint) error
		UpdateSettings(meetingId uint, settings models.AllSettings) error
		AddUserToMeeting(meetingId, userId uint) error
		KickUserFromMeeting(meetingId, userId uint) error
	}

	ParticipationService interface {
		HandleParticipationRequest(request models.ParticipationRequest) (models.RejectInfo, error)
	}

	UsersSettings interface {
		GetUserSettings(userId uint) (models.FullUserInfo, error)
		UpdateUserSettings(userId uint, info models.UserSettings) error
	}

	ChatAccessor interface {
		GetMeetingChat(meetingId uint) (models.Chat, error)
		GetUserChats(userId uint) ([]models.Chat, error)
	}

	Chat interface {
		CreateMeetingChat(meetingId uint) error
		CreateMeetingRequestChat(meetingId uint) error
		CloseChat(chatId uint) error
	}
)
