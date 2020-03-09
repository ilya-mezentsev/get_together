package logging

import (
	"interfaces"
	"models"
	"plugins/logger"
)

type UserSettingsRepositoryDecorator struct {
	repository interfaces.UsersSettingsRepository
}

func NewUserSettingsRepositoryDecorator(
	repository interfaces.UsersSettingsRepository) UserSettingsRepositoryDecorator {
	return UserSettingsRepositoryDecorator{repository}
}

func (d UserSettingsRepositoryDecorator) GetUserSettings(userId uint) (models.FullUserInfo, error) {
	info, err := d.repository.GetUserSettings(userId)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while getting user settings: %v",
			Args: []interface{}{err},
			Optional: map[string]interface{}{
				"user_id": userId,
			},
		}, logger.Warning)
	}

	return info, err
}

func (d UserSettingsRepositoryDecorator) UpdateUserSettings(userId uint, info models.UserSettings) error {
	err := d.repository.UpdateUserSettings(userId, info)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Error while updating user settings: %v",
			Args: []interface{}{err},
			Optional: map[string]interface{}{
				"user_id": userId,
				"user_settings": info,
			},
		}, logger.Warning)
	}

	return err
}

