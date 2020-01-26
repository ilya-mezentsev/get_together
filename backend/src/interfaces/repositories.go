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
  }

  UsersRepository interface {
    GetUserInfo(userId uint) (models.FullUserInfo, error)
    UpdateUserInfo(userId uint, info models.Info) error
    UpdateUserPassword(userId uint, password string) error
  }

  ChatRepository interface {
    CreateMeetingChat(meetingId uint) error
    CreateMeetingRequestChat(meetingId uint) error
    CloseChat(chatId uint) error
  }
)
