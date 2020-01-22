package models

import "time"

type (
  Tag string
  Latitude float64
  Longitude float64
)

type (
  Place struct {
    Label string `db:"label"`
    Latitude `db:"latitude"`
    Longitude `db:"longitude"`
  }

  Settings struct {
    Duration uint `db:"duration"`
    MinAge uint `db:"min_age"`
    Gender string `db:"gender"`
    RequestDescriptionRequired bool `db:"request_description_required"`
  }

  Meeting struct {
    ID uint `db:"id"`
    AdminId uint `db:"admin_id"`
    Title string `db:"title"`
    MaxUsers uint `db:"max_users"`
    Tags []Tag `db:"tags"`
    UserIds []uint `db:"user_ids"`
    DateTime time.Time `db:"date_time"`
    Description string `db:"description"`
    Place
    Settings
  }
)
