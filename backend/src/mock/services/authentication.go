package services

import (
	"errors"
	"internal_errors"
	"mock/repositories"
	"models"
	"utils"
)

type CredentialsMock struct {
	Users []models.UserCredentials
}

var (
	CredentialsRepo = CredentialsMock{
		Users: allUsersCredentials(),
	}
	NewUser = models.UserCredentials{
		Email:    "test@test.ru",
		Password: "pass",
	}
	BadUser = models.UserCredentials{
		Email:    "bad-email@ya.ru",
		Password: "bad password",
	}

	BadUserId         uint = 0
	someInternalError      = errors.New("something went wrong")
)

func (c *CredentialsMock) ResetState() {
	c.Users = allUsersCredentials()
}

func (c *CredentialsMock) CreateUser(user models.UserCredentials) error {
	if c.HasEmail(user.Email) {
		return internal_errors.UnableToRegisterUserEmailExists
	} else if user.Email == BadUser.Email {
		return someInternalError
	}

	c.Users = append(c.Users, user)
	return nil
}

func (c *CredentialsMock) HasEmail(email string) bool {
	for _, user := range c.Users {
		if user.Email == email {
			return true
		}
	}

	return false
}

func (c *CredentialsMock) GetUserIdByCredentials(user models.UserCredentials) (uint, error) {
	if user.Email == BadUser.Email {
		return 0, someInternalError
	}

	for userIdx, registered := range c.Users {
		if registered == user {
			return uint(userIdx) + 1, nil
		}
	}

	return 0, internal_errors.UnableToLoginUserNotFound
}

func (c *CredentialsMock) UpdateUserPassword(user models.UserCredentials) error {
	if user.Email == BadUser.Email {
		return someInternalError
	} else if user.Password == utils.GetHash(user.Email+BadUser.Password) {
		return someInternalError
	}

	for id, u := range c.Users {
		if u.Email == user.Email {
			c.Users[id].Password = user.Password
			return nil
		}
	}

	return internal_errors.UnableToChangePasswordUserNotFound
}

func (c *CredentialsMock) GetUserEmail(userId uint) (string, error) {
	if userId == BadUserId {
		return "", someInternalError
	}

	for id, user := range c.Users {
		if id+1 == int(userId) {
			return user.Email, nil
		}
	}

	return "", internal_errors.UnableToFindUserById
}

func FirstUserCredentials() models.UserCredentials {
	return models.UserCredentials{
		Email:    repositories.UsersCredentials[0]["email"].(string),
		Password: repositories.TestingPassword,
	}
}

func allUsersCredentials() []models.UserCredentials {
	var credentials []models.UserCredentials
	for _, c := range repositories.UsersCredentials {
		credentials = append(credentials, models.UserCredentials{
			Email:    c["email"].(string),
			Password: c["password"].(string),
		})
	}

	return credentials
}
