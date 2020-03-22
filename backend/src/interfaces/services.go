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

	SessionService interface {
		GetSession(r *http.Request) (map[string]interface{}, error)
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
)
