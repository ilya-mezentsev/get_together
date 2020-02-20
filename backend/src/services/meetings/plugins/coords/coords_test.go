package coords

import (
  mock "mock/plugins"
  "models"
  "testing"
  "utils"
)

func sortedByCoords(meetings []models.ExtendedMeeting) bool {
  for i := 0; i < len(meetings)-1; i++ {
    if getVectorLength(meetings[i]) > getVectorLength(meetings[i+1]) {
      return false
    }
  }

  return true
}

func TestShake_Sorting(t *testing.T) {
  meetings := ShakeExtendedMeetings(mock.SortTest)
  utils.AssertTrue(sortedByCoords(meetings), t)
}

func TestShake_Shaking(t *testing.T) {
  meetings := ShakeExtendedMeetings(mock.ShakingTest)

  for _, meeting := range meetings[:10] {
    utils.AssertEqual(7.7, float64(meeting.GetLatitude()), t)
    utils.AssertEqual(10.2,  float64(meeting.GetLongitude()), t)
  }
}
