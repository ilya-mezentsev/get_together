package models

type (
  Rating struct {
    Tag Tag `db:"tag"`
    Value float64 `db:"value"`
  }

  Info struct {
    Name string `db:"name"`
    Nickname string `db:"nickname"`
    Gender string `db:"gender"`
    Age uint `db:"age"`
    AvatarUrl string `db:"avatar_url"`
  }

  AuthUser struct {
    Email string `db:"email"`
    Password string `db:"password"`
  }

  FullUserInfo struct {
    ID uint `db:"id"`
    Info
    Rating
  }
)
