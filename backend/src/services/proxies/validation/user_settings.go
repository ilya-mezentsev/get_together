package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type UserSettingsServiceProxy struct {
	service interfaces.UsersSettings
}

func NewUserSettingsServiceProxy(service interfaces.UsersSettings) UserSettingsServiceProxy {
	return UserSettingsServiceProxy{service}
}

func (p UserSettingsServiceProxy) GetUserSettings(userId uint) (models.FullUserInfo, error) {
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return models.FullUserInfo{}, validationResults
	}

	return p.service.GetUserSettings(userId)
}

func (p UserSettingsServiceProxy) UpdateUserSettings(userId uint, info models.UserSettings) error {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults.Add(InvalidID)
	}
	if !validation.ValidName(info.Name) {
		validationResults.Add(InvalidUserName)
	}
	if !validation.ValidNickname(info.Nickname) {
		validationResults.Add(InvalidUserNickname)
	}
	if !validation.ValidGender(info.Gender) {
		validationResults.Add(InvalidUserGender)
	}
	if !validation.ValidWholePositiveNumber(float64(info.Age)) {
		validationResults.Add(InvalidUserAge)
	}
	if !validation.ValidURL(info.AvatarUrl) {
		validationResults.Add(InvalidUserAvatarURL)
	}

	if validationResults.HasErrors() {
		return validationResults
	} else {
		return p.service.UpdateUserSettings(userId, info)
	}
}



