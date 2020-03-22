package validation

import (
	"interfaces"
	"models"
	"net/http"
	"services/proxies/validation/plugins/validation"
)

type SessionServiceProxy struct {
	service interfaces.SessionService
}

func NewSessionServiceProxy(service interfaces.SessionService) SessionServiceProxy {
	return SessionServiceProxy{service}
}

func (p SessionServiceProxy) GetSession(r *http.Request) (map[string]interface{}, error) {
	return p.service.GetSession(r)
}

func (p SessionServiceProxy) SetSession(r *http.Request, session models.UserSession) error {
	if !validation.ValidWholePositiveNumber(float64(session.ID)) {
		validationResults := validationResults{}
		validationResults.Add(InvalidID)
		return validationResults
	}

	return p.service.SetSession(r, session)
}

func (p SessionServiceProxy) InvalidateSession(r *http.Request) {
	p.service.InvalidateSession(r)
}
