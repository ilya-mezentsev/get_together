package plugins

import (
  "models"
  "time"
)

var (
  TargetMeetingTimeParameters = models.TimeMeetingParameters{
    DateTime: time.Date(2020, 3, 2, 15, 0, 0, 0, time.UTC),
    Duration: 3,
  }
  MeetingsBefore = []models.TimeMeetingParameters{
    {
      DateTime: time.Date(2020, 3, 2, 12, 0, 0, 0, time.UTC),
      Duration: 2,
    },
  }
  MeetingsAfter = []models.TimeMeetingParameters{
    {
      DateTime: time.Date(2020, 3, 2, 18, 0, 0, 0, time.UTC),
      Duration: 2,
    },
  }
  SameTimeMeetings = []models.TimeMeetingParameters{
    TargetMeetingTimeParameters,
  }
  MeetingsLongAfter = []models.TimeMeetingParameters{
    {
      DateTime: time.Date(2020, 3, 2, 21, 0, 0, 0, time.UTC),
      Duration: 2,
    },
  }
  MeetingsLongBefore = []models.TimeMeetingParameters{
    {
      DateTime: time.Date(2020, 3, 2, 10, 0, 0, 0, time.UTC),
      Duration: 2,
    },
  }
)
