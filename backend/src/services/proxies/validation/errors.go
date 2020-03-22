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
	InvalidID                              = "invalid-id"
	InvalidMeetingTitle                    = "invalid-meeting-title"
	InvalidMeetingDescription              = "invalid-meeting-description"
	InvalidMeetingTag                      = "invalid-meeting-tag"
	InvalidMeetingDate                     = "invalid-meeting-date"
	InvalidMeetingMaxUsers                 = "invalid-meeting-max-users"
	InvalidMeetingDuration                 = "invalid-meeting-duration"
	InvalidMeetingMinAge                   = "invalid-meeting-min-age"
	InvalidMeetingGender                   = "invalid-meeting-gender"
	InvalidParticipationRequestDescription = "invalid-participation-request-description"
	InvalidUserName                        = "invalid-user-name"
	InvalidUserNickname                    = "invalid-user-nickname"
	InvalidUserGender                      = "invalid-user-gender"
	InvalidUserAge                         = "invalid-user-age"
	InvalidUserAvatarURL                   = "invalid-user-avatar-url"
)
