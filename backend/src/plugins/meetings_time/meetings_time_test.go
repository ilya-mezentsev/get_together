package meetings_time

import (
  mock "mock/plugins"
  "testing"
  "utils"
)

func TestMeetingsNearTo_HasBefore(t *testing.T) {
  utils.AssertTrue(MeetingsNearTo(mock.TargetMeetingTimeParameters, mock.MeetingsBefore), t)
}

func TestMeetingsNearTo_HasAfter(t *testing.T) {
  utils.AssertTrue(MeetingsNearTo(mock.TargetMeetingTimeParameters, mock.MeetingsAfter), t)
}

func TestMeetingsNearTo_HasSameTime(t *testing.T) {
  utils.AssertTrue(MeetingsNearTo(mock.TargetMeetingTimeParameters, mock.SameTimeMeetings), t)
}

func TestMeetingsNearTo_HasNotAfter(t *testing.T) {
  utils.AssertFalse(MeetingsNearTo(mock.TargetMeetingTimeParameters, mock.MeetingsLongAfter), t)
}

func TestMeetingsNearTo_HasNoBefore(t *testing.T) {
  utils.AssertFalse(MeetingsNearTo(mock.TargetMeetingTimeParameters, mock.MeetingsLongBefore), t)
}
