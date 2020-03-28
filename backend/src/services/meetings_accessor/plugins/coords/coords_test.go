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
	meetings := ShakeExtendedMeetings(mock.GetExtendedMeetingsForShakingTest())

	for _, meeting := range meetings[:10] {
		utils.AssertEqual(9.7, float64(meeting.GetLatitude()), t)
		utils.AssertEqual(12.2, float64(meeting.GetLongitude()), t)
	}
}

func TestShake_ShakingOfOneMeeting(t *testing.T) {
	firstMeeting := mock.GetExtendedMeetingsForShakingTest()[0]
	meetings := ShakeExtendedMeetings(mock.GetExtendedMeetingsForShakingTest()[:1])

	utils.AssertTrue(firstMeeting.GetLongitude() != meetings[0].GetLongitude(), t)
	utils.AssertTrue(firstMeeting.GetLatitude() != meetings[0].GetLatitude(), t)
}

func TestShakePublicMeetings_ShakingOne(t *testing.T) {
	firstMeeting := mock.GetPublicMeetingsForShakingTest()[0]
	meetings := ShakePublicMeetings(mock.GetPublicMeetingsForShakingTest()[:1])

	utils.AssertTrue(firstMeeting.GetLongitude() != meetings[0].GetLongitude(), t)
	utils.AssertTrue(firstMeeting.GetLatitude() != meetings[0].GetLatitude(), t)
}
