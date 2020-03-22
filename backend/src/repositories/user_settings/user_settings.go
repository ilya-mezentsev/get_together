package user_settings

import (
	"github.com/jmoiron/sqlx"
	"internal_errors"
	"models"
)

const (
	GetUserInfoQuery   = `SELECT name, nickname, gender, age, avatar_url FROM users_info WHERE user_id = $1`
	GetUserRatingQuery = `SELECT tag, value FROM users_rating WHERE user_id = $1`
	UpdateInfoQuery    = `
  UPDATE users_info
  SET name = :name, nickname = :nickname, gender = :gender, age = :age, avatar_url = :avatar_url
  WHERE user_id = :user_id`

	userNotFoundMessage = "sql: no rows in result set"
)

type (
	Repository struct {
		db *sqlx.DB
	}

	UpdateInfo struct {
		ID uint `db:"user_id"`
		models.UserSettings
	}
)

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) GetUserSettings(userId uint) (models.FullUserInfo, error) {
	var (
		err          error
		fullUserInfo models.FullUserInfo
	)
	fullUserInfo.UserSettings, err = r.getUserInfo(userId)
	if err != nil {
		return models.FullUserInfo{}, err
	}

	fullUserInfo.Rating, err = r.getUserRating(userId)
	if err != nil {
		return models.FullUserInfo{}, err
	}

	return fullUserInfo, nil
}

func (r Repository) getUserInfo(userId uint) (models.UserSettings, error) {
	var info models.UserSettings
	err := r.db.Get(&info, GetUserInfoQuery, userId)

	switch {
	case err == nil:
		return info, err
	case err.Error() == userNotFoundMessage:
		return models.UserSettings{}, internal_errors.UnableToFindUserById
	default:
		return models.UserSettings{}, err
	}
}

func (r Repository) getUserRating(userId uint) ([]models.Rating, error) {
	var rating []models.Rating
	rows, err := r.db.Queryx(GetUserRatingQuery, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r models.Rating
		if err = rows.StructScan(&r); err != nil {
			return nil, err
		}

		rating = append(rating, r)
	}

	return rating, nil
}

func (r Repository) UpdateUserSettings(userId uint, info models.UserSettings) error {
	res, err := r.db.NamedExec(UpdateInfoQuery, UpdateInfo{
		ID:           userId,
		UserSettings: info,
	})
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affectedRows == 0 {
		return internal_errors.UnableToFindUserById
	}

	return nil
}
