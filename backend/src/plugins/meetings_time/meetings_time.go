package meetings_time

import (
  "models"
  "time"
)

const (
  nanosecondsInHour = 36e11
  defaultHoursDelta = 2 * nanosecondsInHour
)

func MeetingsNearTo(checkingMeeting models.TimeMeetingParameters, meetings []models.TimeMeetingParameters) bool {
  for _, meeting := range meetings {
    if checkingMeeting.DateTime.Before(meeting.DateTime) && meetingNearTo(checkingMeeting, meeting) {
      return true
    } else if checkingMeeting.DateTime.After(meeting.DateTime) && meetingNearTo(meeting, checkingMeeting) {
      return true
    } else if checkingMeeting.DateTime.Equal(meeting.DateTime) {
      return true
    }
  }

  return false
}

func meetingNearTo(beforeMeeting, afterMeeting models.TimeMeetingParameters) bool {
  dateAfterDurationAndDelta :=
    beforeMeeting.DateTime.Add(
      time.Duration(beforeMeeting.Duration * nanosecondsInHour)).Add(
        time.Duration(defaultHoursDelta))
  return dateAfterDurationAndDelta == afterMeeting.DateTime || dateAfterDurationAndDelta.After(afterMeeting.DateTime)
}
