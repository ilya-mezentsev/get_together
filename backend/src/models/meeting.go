package models

import "time"

type (
  Tag string
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
    Tags []Tag `db:"tags"`
  }

  // for authorized users
  ExtendedSettings struct {
    PublicSettings
    DateTime time.Time `db:"date_time"`
    RequestDescriptionRequired bool `db:"request_description_required"`
  }

  AllSettings struct {
    ExtendedSettings
    LabeledPlace
    Duration uint `db:"duration"`
    MinAge uint `db:"min_age"`
    Gender string `db:"gender"`
    MaxUsers uint `db:"max_users"`
  }

  DefaultMeeting struct {
    ID uint `db:"id"`
    AdminId uint `db:"admin_id"`
    CreatedAt time.Time `db:"created_at"`
  }

  PublicMeeting struct {
    DefaultMeeting
    PublicSettings
    PublicPlace
  }

  ExtendedMeeting struct {
    DefaultMeeting
    ExtendedSettings
    PublicPlace
    CurrentUserStatus string
  }

  // for invited users
  PrivateMeeting struct {
    DefaultMeeting
    LabeledPlace
    AllSettings
  }
)
