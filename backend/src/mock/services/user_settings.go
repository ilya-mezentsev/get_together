package services

import (
	"internal_errors"
	"mock/repositories"
	"models"
)

type UsersSettingsRepositoryMock struct {
	Settings map[uint]models.FullUserInfo
}

var (
	UsersSettingsRepository = UsersSettingsRepositoryMock{Settings: allUsersSettings()}
	NewUserInfo             = models.UserSettings{
		Name:   "Hello world",
		Age:    16,
		Gender: "male",
	}
)

func (u *UsersSettingsRepositoryMock) ResetState() {
	u.Settings = allUsersSettings()
}

func (u *UsersSettingsRepositoryMock) GetUserSettings(userId uint) (models.FullUserInfo, error) {
	if userId == BadUserId {
		return models.FullUserInfo{}, someInternalError
	}

	userInfo, found := u.Settings[userId]
	if !found {
		return models.FullUserInfo{}, internal_errors.UnableToFindUserById
	}

	return userInfo, nil
}

func (u *UsersSettingsRepositoryMock) UpdateUserSettings(userId uint, info models.UserSettings) error {
	if userId == BadUserId {
		return someInternalError
	}

	_, found := u.Settings[userId]
	if !found {
		return internal_errors.UnableToFindUserById
	}

	u.Settings[userId] = models.FullUserInfo{
		UserSettings: info,
		Rating:       nil,
	}
	return nil
}

func allUsersSettings() map[uint]models.FullUserInfo {
	settings := map[uint]models.FullUserInfo{}
	for _, u := range repositories.UsersInfo {
		userId := uint(u["user_id"].(int))
		settings[userId] = models.FullUserInfo{
			UserSettings: models.UserSettings{
				Name:     u["name"].(string),
				Nickname: u["nickname"].(string),
				Gender:   u["gender"].(string),
				Age:      uint(u["age"].(int)),
			},
			Rating: getRatingByUserId(userId),
		}
	}

	return settings
}

func getRatingByUserId(userId uint) []models.Rating {
	var rating []models.Rating
	for _, r := range repositories.UsersRating {
		if uint(r["user_id"].(int)) != userId {
			continue
		}

		rating = append(rating, models.Rating{
			Tag:   r["tag"].(string),
			Value: float64(r["value"].(int)),
		})
	}

	return rating
}
