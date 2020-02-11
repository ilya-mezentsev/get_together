package models

type (
  ParticipationRequest struct {
    UserId uint `json:"user_id"`
    MeetingId uint `json:"meeting_id"`
    RequestDescription string `json:"request_description"`
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
