package models

type (
	Rating struct {
		Tag   string  `db:"tag" json:"tag"`
		Value float64 `db:"value" json:"value"`
	}

	UserSettings struct {
		Name      string `db:"name" json:"name"`
		Nickname  string `db:"nickname" json:"nickname"`
		Gender    string `db:"gender" json:"gender"`
		Age       uint   `db:"age" json:"age"`
		AvatarUrl string `db:"avatar_url" json:"avatar_url"`
	}

	UserCredentials struct {
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
	}

	FullUserInfo struct {
		UserSettings
		Rating []Rating `json:"rating"`
	}
)

type (
	UserRatingLevelData struct {
		UserId      uint `db:"user_id"`
		MeetingId   uint `db:"meeting_id"`
		BottomValue uint `db:"bottom_value"`
	}

	UserTimeCheckData struct {
		UserId    uint `db:"user_id"`
		MeetingId uint `db:"meeting_id"`
	}
)
