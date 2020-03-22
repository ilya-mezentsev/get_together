package repositories

import (
	"models"
)

var (
	NotExistsUser = models.UserCredentials{
		Email:    "no-way@ya.ru",
		Password: "",
	}
	NewUser = models.UserCredentials{
		Email:    "new@mail.ru",
		Password: "some_pass",
	}
)

func GetFirstUser() models.UserCredentials {
	return models.UserCredentials{
		Email:    UsersCredentials[0]["email"].(string),
		Password: UsersCredentials[0]["password"].(string),
	}
}

func GetNextUserId() uint {
	return uint(len(UsersCredentials) + 1)
}
