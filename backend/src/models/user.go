package models

type (
  Rating struct {
    Value float64 `db:"current_value"`
  }

  Info struct {
    Name string `db:"name"`
    Nickname string `db:"nickname"`
    Gender string `db:"gender"`
    Age uint `db:"age"`
    AvatarUrl string `db:"avatar_url"`
  }

  User struct {
    ID uint `db:"id"`
    Email string `db:"email"`
    Password string `db:"password"`
    Info
    Rating
  }
)
