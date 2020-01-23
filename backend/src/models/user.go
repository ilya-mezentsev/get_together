package models

type (
  Rating struct {
    Value float64
  }

  Info struct {
    Name string `db:"name"`
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
