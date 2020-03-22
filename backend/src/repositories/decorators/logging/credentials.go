package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type CredentialsRepositoryDecorator struct {
	repository interfaces.CredentialsRepository
}

func NewCredentialsRepositoryDecorator(repository interfaces.CredentialsRepository) CredentialsRepositoryDecorator {
	return CredentialsRepositoryDecorator{repository}
}

func (d CredentialsRepositoryDecorator) CreateUser(user models.UserCredentials) error {
	err := d.repository.CreateUser(user)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while creating user: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"credentials": user,
			},
		}, logger.Warning)
	}

	return err
}

func (d CredentialsRepositoryDecorator) GetUserIdByCredentials(user models.UserCredentials) (uint, error) {
	userId, err := d.repository.GetUserIdByCredentials(user)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting user id by credentials: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"credentials": user,
			},
		}, logger.Warning)
	}

	return userId, err
}

func (d CredentialsRepositoryDecorator) UpdateUserPassword(user models.UserCredentials) error {
	err := d.repository.UpdateUserPassword(user)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while updating user password: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"credentials": user,
			},
		}, logger.Warning)
	}

	return err
}

func (d CredentialsRepositoryDecorator) GetUserEmail(userId uint) (string, error) {
	email, err := d.repository.GetUserEmail(userId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting user email: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"user_id": userId,
			},
		}, logger.Warning)
	}

	return email, err
}
