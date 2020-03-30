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
		CreateChat(meetingId uint, chatType string) error
		SetChatStatus(chatId uint, status string) error
	}

	FullChatsRepository interface {
		ChatAccessor
		ChatRepository
	}

	FullMeetingsRepository interface {
		Meetings
		MeetingsAccessorRepository
	}
)
