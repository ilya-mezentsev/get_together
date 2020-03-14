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

	MeetingsService interface {
		GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error)
		GetPublicMeetings() ([]models.PublicMeeting, error)
		GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error)
		CreateMeeting(adminId uint, settings models.AllSettings) error
		DeleteMeeting(meetingId uint) error
		UpdatedSettings(meetingId uint, settings models.AllSettings) error
	}

	ParticipationService interface {
		HandleParticipationRequest(request models.ParticipationRequest) (models.RejectInfo, error)
	}

	UserSettingsService interface {
		GetUserSettings(userId uint) (models.FullUserInfo, error)
		UpdateUserSettings(userId uint, info models.UserSettings) error
	}
)
