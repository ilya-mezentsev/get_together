package meetings

import (
	"api"
	"github.com/gorilla/mux"
	"interfaces"
	"models"
	"net/http"
	"strconv"
)

type Handler struct {
	meetingsService         interfaces.Meetings
	participationService    interfaces.ParticipationService
	meetingsAccessorService interfaces.MeetingsAccessorService
}

func InitRequestHandlers(
	meetingsService interfaces.Meetings,
	participationService interfaces.ParticipationService,
	meetingsAccessorService interfaces.MeetingsAccessorService,
) {
	handler := Handler{
		meetingsService,
		participationService,
		meetingsAccessorService,
	}

	meetingsAPI := api.GetRouter().PathPrefix("/meeting").Subrouter()

	api.GetRouter().HandleFunc("/meetings", handler.getPublicMeetings).Methods(http.MethodGet)
	api.GetRouter().HandleFunc("/meetings/{id:[0-9]+}", handler.getExtendedMeetings).Methods(http.MethodGet)
	meetingsAPI.HandleFunc("/", handler.createMeeting).Methods(http.MethodPost)
	meetingsAPI.HandleFunc("/", handler.deleteMeeting).Methods(http.MethodDelete)
	meetingsAPI.HandleFunc("/settings", handler.updateMeetingSettings).Methods(http.MethodPatch)
	meetingsAPI.HandleFunc("/request-participation", handler.handleParticipationRequest).Methods(http.MethodPost)
	meetingsAPI.HandleFunc("/user", handler.inviteUser).Methods(http.MethodPost)
	meetingsAPI.HandleFunc("/user", handler.kickUser).Methods(http.MethodDelete)
}

func (h Handler) getPublicMeetings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	meetings, err := h.meetingsAccessorService.GetPublicMeetings()
	if err != nil {
		panic(err)
	}

	api.EncodeAndSendResponse(w, meetings)
}

func (h Handler) getExtendedMeetings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	vars := mux.Vars(r)
	// checking of this parameter will be performed in validation proxy
	userId, _ := strconv.Atoi(vars["id"])
	meetings, err := h.meetingsAccessorService.GetExtendedMeetings(uint(userId))
	if err != nil {
		panic(err)
	}

	api.EncodeAndSendResponse(w, meetings)
}

func (h Handler) createMeeting(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.CreateMeetingRequest
	api.DecodeRequestBody(r, &request)

	err := h.meetingsService.CreateMeeting(request.AdminId, request.Settings)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}

func (h Handler) deleteMeeting(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.GeneralMeetingRequest
	api.DecodeRequestBody(r, &request)

	err := h.meetingsService.DeleteMeeting(request.MeetingId)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}

func (h Handler) updateMeetingSettings(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.UpdateMeetingSettingsRequest
	api.DecodeRequestBody(r, &request)

	err := h.meetingsService.UpdateSettings(request.MeetingId, request.Settings)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}

func (h Handler) handleParticipationRequest(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.ParticipationRequest
	api.DecodeRequestBody(r, &request)

	rejectInfo, err := h.participationService.HandleParticipationRequest(request)
	if err != nil {
		panic(err)
	}

	api.EncodeAndSendResponse(w, rejectInfo)
}

func (h Handler) inviteUser(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.MeetingUserRequest
	api.DecodeRequestBody(r, &request)

	err := h.meetingsService.AddUserToMeeting(request.MeetingId, request.UserId)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}

func (h Handler) kickUser(w http.ResponseWriter, r *http.Request) {
	defer api.SendErrorIfPanicked(w)

	var request models.MeetingUserRequest
	api.DecodeRequestBody(r, &request)

	err := h.meetingsService.KickUserFromMeeting(request.MeetingId, request.UserId)
	if err != nil {
		panic(err)
	}

	api.SendDefaultResponse(w)
}
