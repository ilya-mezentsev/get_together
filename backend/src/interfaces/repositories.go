package interfaces

import "models"

type (
	MeetingsAccessorRepository interface {
		GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error)
		GetPublicMeetings() ([]models.PublicMeeting, error)
		GetExtendedMeetings(userStatusesData models.UserMeetingStatusesData) ([]models.ExtendedMeeting, error)
	}

	MeetingsSettingsRepository interface {
		GetMeetingSettings(meetingId uint) (models.ParticipationMeetingSettings, error)
		GetNearMeetings(data models.UserTimeCheckData) ([]models.TimeMeetingParameters, error)
	}

	CredentialsRepository interface {
		CreateUser(user models.UserCredentials) error
		GetUserIdByCredentials(user models.UserCredentials) (uint, error)
		UpdateUserPassword(user models.UserCredentials) error
		GetUserEmail(userId uint) (string, error)
	}

	ChatRepository interface {
		CreateMeetingChat(meetingId uint) error
		CreateMeetingRequestChat(meetingId uint) error
		CloseChat(chatId uint) error
	}

	FullMeetingsRepository interface {
		Meetings
		MeetingsAccessorRepository
	}
)
