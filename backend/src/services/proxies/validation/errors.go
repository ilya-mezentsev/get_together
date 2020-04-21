package validation

import "strings"

type validationResults struct {
	validation []string
}

func (e *validationResults) Add(err string) {
	e.validation = append(e.validation, err)
}

func (e validationResults) HasErrors() bool {
	return len(e.validation) != 0
}

func (e validationResults) Error() string {
	return strings.Join(e.validation, "|")
}

const (
	InvalidEmail                           = "invalid-email"
	InvalidPassword                        = "invalid-password"
	InvalidId                              = "invalid-id"
	InvalidMeetingTitle                    = "invalid-meeting-title"
	InvalidMeetingDescription              = "invalid-meeting-description"
	InvalidMeetingTag                      = "invalid-meeting-tag"
	InvalidMeetingMaxUsers                 = "invalid-meeting-max-users"
	InvalidMeetingDuration                 = "invalid-meeting-duration"
	InvalidMeetingMinAge                   = "invalid-meeting-min-age"
	InvalidMeetingGender                   = "invalid-meeting-gender"
	InvalidMeetingLatitude                 = "invalid-meeting-latitude"
	InvalidMeetingLongitude                = "invalid-meeting-longitude"
	InvalidMeetingLabel                    = "invalid-meeting-label"
	InvalidParticipationRequestDescription = "invalid-participation-request-description"
	InvalidUserName                        = "invalid-user-name"
	InvalidUserNickname                    = "invalid-user-nickname"
	InvalidUserGender                      = "invalid-user-gender"
	InvalidUserAge                         = "invalid-user-age"
	InvalidUserAvatarURL                   = "invalid-user-avatar-url"
	InvalidCount                           = "invalid-count"
	InvalidDate                            = "invalid-date"
	InvalidMessageText                     = "invalid-message-text"
)
