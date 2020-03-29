package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"mock/repositories"
	"models"
	"net/http"
	"utils"
)

const (
	testMeetingSettings = `{
		"title": "meeting title",
		"date_time": "02-01-2020 15:00:00",
		"label": "address of meeting",
		"max_users": 10,
		"tags": ["tag1", "tag2"],
		"description": "some meeting description",
		"duration": 3,
		"min_age": 16,
		"male": "male",
		"request_description_required": true
	}`
)

type (
	PublicMeetingsResponse struct {
		Status string                 `json:"status"`
		Data   []models.PublicMeeting `json:"data"`
	}

	ExtendedMeetingsResponse struct {
		Status string                   `json:"status"`
		Data   []models.ExtendedMeeting `json:"data"`
	}

	HandleParticipationResponse struct {
		Status string            `json:"status"`
		Data   models.RejectInfo `json:"data"`
	}
)

func GetPublicMeetingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "meetings",
		Cookie:   &http.Cookie{},
	}
}

func GetExtendedMeetingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "meetings/1",
		Cookie:   &http.Cookie{},
	}
}

func CreateMeetingRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     getNewMeetingSettings(1),
	}
}

func getNewMeetingSettings(adminId uint) string {
	return fmt.Sprintf(`{
			"admin_id": %d,
			"settings": %s
		}`, adminId, testMeetingSettings)
}

func CreateMeetingInvalidAdminIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     getNewMeetingSettings(0),
	}
}

func CreateMeetingNotExistsAdminIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     getNewMeetingSettings(uint(len(repositories.UsersCredentials) + 1)),
	}
}

func CreateMeetingWithInvalidSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     `{"admin_id": 1, "settings": {"latitude": 91, "longitude": -181}}`,
	}
}

func DeleteMeetingRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     `{"meeting_id": 1}`,
	}
}

func DeleteMeetingIDNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     fmt.Sprintf(`{"meeting_id": %d}`, len(repositories.Meetings)+1),
	}
}

func DeleteMeetingByInvalidIDRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   &http.Cookie{},
		Data:     `{"meeting_id": 0}`,
	}
}

func UpdateMeetingSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   &http.Cookie{},
		Data:     getUpdateMeetingSettingsRequestData(1),
	}
}

func getUpdateMeetingSettingsRequestData(meetingId uint) string {
	return fmt.Sprintf(`{"meeting_id": %d, "settings": %s}`, meetingId, testMeetingSettings)
}

func UpdateMeetingSettingsMeetingIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   &http.Cookie{},
		Data:     getUpdateMeetingSettingsRequestData(uint(len(repositories.Meetings) + 1)),
	}
}

func UpdateMeetingSettingsByInvalidMeetingIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   &http.Cookie{},
		Data:     getUpdateMeetingSettingsRequestData(0),
	}
}

func ParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(2, 1),
	}
}

func UserIDNotFoundParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(2, uint(len(repositories.UsersCredentials)+1)),
	}
}

func MeetingIDNotFoundParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   &http.Cookie{},
		Data:     getMeetingIDNotFoundUserRequestData(),
	}
}

func InvalidIDsParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserInvalidIDsRequestData(),
	}
}

func InviteUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(1, 2),
	}
}

func InviteAlreadyInMeetingUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(1, 1),
	}
}

func InviteUserMeetingIDNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingIDNotFoundUserRequestData(),
	}
}

func InviteUserMeetingInvalidIDsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserInvalidIDsRequestData(),
	}
}

func KickUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(1, 1),
	}
}

func KickNotInMeetingUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserRequestData(1, 2),
	}
}

func KickUserInvalidIDsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   &http.Cookie{},
		Data:     getMeetingUserInvalidIDsRequestData(),
	}
}

func getMeetingIDNotFoundUserRequestData() string {
	return getMeetingUserRequestData(uint(len(repositories.Meetings)+1), 1)
}

func getMeetingUserInvalidIDsRequestData() string {
	return getMeetingUserRequestData(0, 0)
}

func getMeetingUserRequestData(meetingId, userId uint) string {
	return fmt.Sprintf(`{"user_id": %d, "meeting_id": %d}`, userId, meetingId)
}
