package validation

import (
	"interfaces"
	"models"
	"services/proxies/validation/plugins/validation"
)

type ParticipationServiceProxy struct {
	service interfaces.ParticipationService
}

func NewParticipationServiceProxy(service interfaces.ParticipationService) ParticipationServiceProxy {
	return ParticipationServiceProxy{service}
}

func (p ParticipationServiceProxy) HandleParticipationRequest(
	request models.ParticipationRequest) (models.RejectInfo, error) {
	validationResults := validationResults{}
	if !validation.ValidWholePositiveNumber(float64(request.UserId)) ||
		!validation.ValidWholePositiveNumber(float64(request.MeetingId)) {
		validationResults.Add(InvalidID)
	}
	if request.RequestDescription != "" && !validation.ValidDescription(request.RequestDescription) {
		validationResults.Add(InvalidParticipationRequestDescription)
	}

	if validationResults.HasErrors() {
		return models.RejectInfo{}, validationResults
	} else {
		return p.service.HandleParticipationRequest(request)
	}
}



