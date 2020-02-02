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
  meetings := Shake(mock.SortTest)
  utils.Assert(sortedByCoords(meetings), func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "sorted by coords meetings", Got: meetings}))
    t.Fail()
  })
}

func TestShake_Shaking(t *testing.T) {
  meetings := Shake(mock.ShakingTest)

  for _, meeting := range meetings[:10] {
    utils.Assert(7.7 == meeting.Latitude, func() {
      t.Log(
        utils.GetExpectationString(
          utils.Expectation{Expected: 7.7, Got: meeting.Latitude}))
      t.Fail()
    })
    utils.Assert(10.2 == meeting.Longitude, func() {
      t.Log(
        utils.GetExpectationString(
          utils.Expectation{Expected: 10.2, Got: meeting.Longitude}))
      t.Fail()
    })
  }
}
