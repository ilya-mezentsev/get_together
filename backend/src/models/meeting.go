package models

import "time"

type (
  Latitude float64
  Longitude float64
)

type (
  PublicPlace struct {
    Latitude `db:"latitude"`
    Longitude `db:"longitude"`
  }

  LabeledPlace struct {
    Label string `db:"label"`
    PublicPlace
  }

  // for unauthorized users
  PublicSettings struct {
    Title string `db:"title"`
    Description string `db:"description"`
    Tags []string `db:"tags"`
  }

  // for authorized users
  ExtendedSettings struct {
    PublicSettings
    MeetingParameters
  }

  AllSettings struct {
    ExtendedSettings
    LabeledPlace
    MeetingLimitations
  }

  MeetingParameters struct {
    DateTime time.Time `db:"date_time"`
    RequestDescriptionRequired bool `db:"request_description_required"`
  }

  MeetingLimitations struct {
    MaxUsers uint `db:"max_users"`
    Duration uint `db:"duration"`
    MinAge uint `db:"min_age"`
    Gender string `db:"gender"`
  }

  ParticipationMeetingSettings struct {
    MeetingLimitations
    MeetingParameters
    Tags []string `db:"tags"`
    UsersCount uint `db:"users_count"`
  }

  DefaultMeeting struct {
    ID uint `db:"id"`
    AdminId uint `db:"admin_id"`
    CreatedAt time.Time `db:"created_at"`
  }

  PublicMeeting struct {
    DefaultMeeting
    PublicSettings
    *PublicPlace
  }

  ExtendedMeeting struct {
    DefaultMeeting
    ExtendedSettings
    *PublicPlace
    CurrentUserStatus string `db:"current_user_status"`
  }

  // for invited users
  PrivateMeeting struct {
    DefaultMeeting
    *LabeledPlace
    AllSettings
  }
)

type (
  UserMeetingStatusesData struct {
    UserId uint `db:"user_id"`
    Invited string `db:"invited"`
    NotInvited string `db:"not_invited"`
  }
  TimeMeetingParameters struct {
    DateTime time.Time `db:"date_time"`
    Duration uint `db:"duration"`
  }
)

func (p *PublicPlace) GetLatitude() Latitude {
  return p.Latitude
}

func (p *PublicPlace) SetLatitude(l Latitude) {
  p.Latitude = l
}

func (p *PublicPlace) GetLongitude() Longitude {
  return p.Longitude
}

func (p *PublicPlace) SetLongitude(l Longitude) {
  p.Longitude = l
}
