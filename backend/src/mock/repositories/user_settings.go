package repositories

import (
	"models"
)

var (
	TestInfo = models.UserSettings{
		Nickname: "some_nickname",
		Gender:   "male",
	}
)

func SettingsEqual(s1, s2 models.FullUserInfo) bool {
	if s1.UserSettings != s2.UserSettings {
		return false
	}

	if len(s1.Rating) != len(s2.Rating) {
		return false
	}

	for i := range s1.Rating {
		if s1.Rating[i] != s2.Rating[i] {
			return false
		}
	}

	return true
}

func GetFirstUserSettings() models.FullUserInfo {
	return models.FullUserInfo{
		UserSettings: models.UserSettings{
			Name:     UsersInfo[0]["name"].(string),
			Nickname: UsersInfo[0]["nickname"].(string),
			Gender:   UsersInfo[0]["gender"].(string),
			Age:      uint(UsersInfo[0]["age"].(int)),
		},
		Rating: []models.Rating{
			{Tag: UsersRating[0]["tag"].(string), Value: float64(UsersRating[0]["value"].(int))},
			{Tag: UsersRating[1]["tag"].(string), Value: float64(UsersRating[1]["value"].(int))},
			{Tag: UsersRating[2]["tag"].(string), Value: float64(UsersRating[2]["value"].(int))},
		},
	}
}

func GetNotExistsUserId() uint {
	return uint(len(UsersCredentials) + 1)
}
