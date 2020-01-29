package interfaces

import "models"

type (
  MeetingsRepository interface {
    GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error)
    GetPublicMeetings() ([]models.PublicMeeting, error)
    GetExtendedMeetings(userId uint) ([]models.ExtendedMeeting, error)
    CreateMeeting(adminId uint, settings models.AllSettings) error
    DeleteMeeting(meetingId uint) error
    UpdatedSettings(meetingId uint, settings models.AllSettings) error
    AddUserToMeeting(meetingId, userId uint) error
    KickUserFromMeeting(meetingId, userId uint)
  }

  CredentialsRepository interface {
    CreateUser(user models.UserCredentials) error
    GetUserIdByCredentials(user models.UserCredentials) (uint, error)
    UpdateUserPassword(user models.UserCredentials) error
    GetUserEmail(userId uint) (string, error)
  }

  UsersSettingsRepository interface {
    GetUserSettings(userId uint) (models.FullUserInfo, error)
    UpdateUserSettings(userId uint, info models.UserSettings) error
  }

  ChatRepository interface {
    CreateMeetingChat(meetingId uint) error
    CreateMeetingRequestChat(meetingId uint) error
    CloseChat(chatId uint) error
  }
)
