package models

type (
  ErrorResponse struct {
    Status string `json:"status"`
    ErrorDetail string `json:"error_detail"`
  }

  SuccessResponse struct {
    Status string `json:"status"`
    Data interface{} `json:"data"`
  }

  UpdateUserSettingsRequest struct {
    UserId uint `json:"user_id"`
    Settings UserSettings `json:"settings"`
  }

  ParticipationRequest struct {
    UserId uint `json:"user_id"`
    MeetingId uint `json:"meeting_id"`
    RequestDescription string `json:"request_description"`
  }

  ChangePasswordRequest struct {
    UserId uint `json:"user_id"`
    Password string `json:"password"`
  }

  InappropriateInfoField struct {
    ErrorCode string `json:"error_code"`
    Description string `json:"description"`
  }

  RejectInfo struct {
    TooLowRatingTags []string `json:"too_low_rating_tags"`
    InappropriateInfoFields []InappropriateInfoField `json:"inappropriate_info_fields"`
    HasNearMeeting bool `json:"has_near_meeting"`
  }
)
