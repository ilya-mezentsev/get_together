package models

type (
  Rating struct {
    Tag string `db:"tag"`
    Value float64 `db:"value"`
  }

  UserSettings struct {
    Name string `db:"name"`
    Nickname string `db:"nickname"`
    Gender string `db:"gender"`
    Age uint `db:"age"`
    AvatarUrl string `db:"avatar_url"`
  }

  UserCredentials struct {
    Email string `db:"email"`
    Password string `db:"password"`
  }

  FullUserInfo struct {
    UserSettings
    Rating []Rating
  }
)

type (
  UserRatingLevelData struct {
    UserId uint `db:"user_id"`
    MeetingId uint `db:"meeting_id"`
    BottomValue uint `db:"bottom_value"`
  }

  UserTimeCheckData struct {
    UserId uint `db:"user_id"`
    MeetingId uint `db:"meeting_id"`
  }
)
