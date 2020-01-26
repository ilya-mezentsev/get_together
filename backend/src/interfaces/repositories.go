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

  UsersRepository interface {
    CreateUser(user models.AuthUser) error
    GetUserInfo(userId uint) (models.FullUserInfo, error)
    GetUserIdByCredentials(user models.AuthUser) (uint, error)
    UpdateUserInfo(userId uint, info models.Info) error
    UpdateUserPassword(userId uint, password string) error
  }

  ChatRepository interface {
    CreateMeetingChat(meetingId uint) error
    CreateMeetingRequestChat(meetingId uint) error
    CloseChat(chatId uint) error
  }
)
