package errors

import "errors"

var (
	InternalError        = errors.New("internal-error")
	UserIdNotFound       = errors.New("user-id-not-found")
	UserAlreadyInMeeting = errors.New("user-already-in-meeting")
	UserNotInMeeting     = errors.New("user-not-in-meeting")
	MeetingIdNotFound    = errors.New("meeting-id-not-found")
	EmailExists          = errors.New("email-exists")
	CredentialsNotFound  = errors.New("credentials-not-found")
	NoAuthCookie         = errors.New("no-auth-cookie")
	InvalidAuthCookie    = errors.New("invalid-auth-cookie")
)
