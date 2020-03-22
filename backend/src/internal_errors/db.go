package internal_errors

import "errors"

var (
	UnableToRegisterUserEmailExists    = errors.New("unable to register user: email already exists")
	UnableToLoginUserNotFound          = errors.New("unable to login user: not found by credentials")
	UnableToChangePasswordUserNotFound = errors.New("unable to change user password: not found by email")
	UnableToFindUserById               = errors.New("unable to find user by id")
	UnableToFindByMeetingId            = errors.New("unable to find meeting by id")
	UserAlreadyInMeeting               = errors.New("user already in meeting")
	UserNotInMeeting                   = errors.New("user not in meeting")
)
