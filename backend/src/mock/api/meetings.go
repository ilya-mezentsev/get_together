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
		Cookie:   emptyCookie,
	}
}

func GetExtendedMeetingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "meetings/1",
		Cookie:   cookie,
	}
}

func GetExtendedMeetingsRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodGet,
		Endpoint: "meetings/1",
		Cookie:   emptyCookie,
	}
}

func CreateMeetingRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     getNewMeetingSettings(1),
	}
}

func CreateMeetingRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   emptyCookie,
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
		Cookie:   cookie,
		Data:     getNewMeetingSettings(0),
	}
}

func CreateMeetingNotExistsAdminIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     getNewMeetingSettings(uint(len(repositories.UsersCredentials) + 1)),
	}
}

func CreateMeetingWithInvalidSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     `{"admin_id": 1, "settings": {"latitude": 91, "longitude": -181}}`,
	}
}

func DeleteMeetingRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     `{"meeting_id": 1}`,
	}
}

func DeleteMeetingRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   emptyCookie,
		Data:     `{"meeting_id": 1}`,
	}
}

func DeleteMeetingIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     fmt.Sprintf(`{"meeting_id": %d}`, len(repositories.Meetings)+1),
	}
}

func DeleteMeetingByInvalidIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/",
		Cookie:   cookie,
		Data:     `{"meeting_id": 0}`,
	}
}

func UpdateMeetingSettingsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   cookie,
		Data:     getUpdateMeetingSettingsRequestData(1),
	}
}

func UpdateMeetingSettingsRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   emptyCookie,
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
		Cookie:   cookie,
		Data:     getUpdateMeetingSettingsRequestData(uint(len(repositories.Meetings) + 1)),
	}
}

func UpdateMeetingSettingsByInvalidMeetingIdRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPatch,
		Endpoint: "meeting/settings",
		Cookie:   cookie,
		Data:     getUpdateMeetingSettingsRequestData(0),
	}
}

func ParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(2, 1),
	}
}

func ParticipationRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   emptyCookie,
		Data:     getMeetingUserRequestData(2, 1),
	}
}

func UserIdNotFoundParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(2, uint(len(repositories.UsersCredentials)+1)),
	}
}

func MeetingIdNotFoundParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   cookie,
		Data:     getMeetingIdNotFoundUserRequestData(),
	}
}

func InvalidIdsParticipationRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/request-participation",
		Cookie:   cookie,
		Data:     getMeetingUserInvalidIdsRequestData(),
	}
}

func InviteUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(1, 2),
	}
}

func InviteUserRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   emptyCookie,
		Data:     getMeetingUserRequestData(1, 2),
	}
}

func InviteAlreadyInMeetingUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(1, 1),
	}
}

func InviteUserMeetingIdNotFoundRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingIdNotFoundUserRequestData(),
	}
}

func InviteUserMeetingInvalidIdsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodPost,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserInvalidIdsRequestData(),
	}
}

func KickUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(1, 1),
	}
}

func KickUserRequestWithoutSession(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   emptyCookie,
		Data:     getMeetingUserRequestData(1, 1),
	}
}

func KickNotInMeetingUserRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserRequestData(1, 2),
	}
}

func KickUserInvalidIdsRequest(r *mux.Router) utils.RequestData {
	return utils.RequestData{
		Router:   r,
		Method:   http.MethodDelete,
		Endpoint: "meeting/user",
		Cookie:   cookie,
		Data:     getMeetingUserInvalidIdsRequestData(),
	}
}

func getMeetingIdNotFoundUserRequestData() string {
	return getMeetingUserRequestData(uint(len(repositories.Meetings)+1), 1)
}

func getMeetingUserInvalidIdsRequestData() string {
	return getMeetingUserRequestData(0, 0)
}

func getMeetingUserRequestData(meetingId, userId uint) string {
	return fmt.Sprintf(`{"user_id": %d, "meeting_id": %d}`, userId, meetingId)
}
