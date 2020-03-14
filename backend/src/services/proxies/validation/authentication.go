package validation

import (
	"models"
	"services/authentication"
	"services/proxies/validation/plugins/validation"
)

type AuthenticationServiceProxy struct {
	service authentication.Service
}

func NewAuthenticationServiceProxy(service authentication.Service) AuthenticationServiceProxy {
	return AuthenticationServiceProxy{service}
}

func (p AuthenticationServiceProxy) RegisterUser(credentials models.UserCredentials) error {
	validationResults := p.validateCredentials(credentials)

	if validationResults.HasErrors() {
		return validationResults
	} else {
		return p.service.RegisterUser(credentials)
	}
}

func (p AuthenticationServiceProxy) validateCredentials(credentials models.UserCredentials) validationResults {
	validationResults := validationResults{}
	if !validation.ValidEmail(credentials.Email) {
		validationResults.Add(InvalidEmail)
	}
	if !validation.ValidPassword(credentials.Password) {
		validationResults.Add(InvalidPassword)
	}

	return validationResults
}

func (p AuthenticationServiceProxy) Login(credentials models.UserCredentials) (models.UserSession, error) {
	validationResults := p.validateCredentials(credentials)

	if len(validationResults.validation) != 0 {
		return models.UserSession{}, validationResults
	} else {
		return p.service.Login(credentials)
	}
}

func (p AuthenticationServiceProxy) ChangePassword(userId uint, password string) error {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(userId)) {
		validationResults.Add(InvalidID)
	}
	if !validation.ValidPassword(password) {
		validationResults.Add(InvalidPassword)
	}

	if len(validationResults.validation) != 0 {
		return validationResults
	} else {
		return p.service.ChangePassword(userId, password)
	}
}
