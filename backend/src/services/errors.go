package services

import "errors"

var (
  InternalError = errors.New("internal-error")
  UserIdNotFound = errors.New("user-id-not-found")
  MeetingIdNotFound = errors.New("meeting-id-not-found")
)
