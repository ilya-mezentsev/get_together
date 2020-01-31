package coords

// General idea:
//  1. Sorting all meetings by length of vector OA, where O(0; 0), A(place.Latitude; place.Longitude)
//  2. Creating meetings blocks (blocks by step == blockSize)
//  3. Set for all meetings in block average coordinates

import (
  "math"
  "models"
)

const blockSize = 10

type Border struct {
  left, right int
}

func Shake(meetings []models.ExtendedMeeting) []models.ExtendedMeeting {
  meetings = sortByCoords(meetings)
  meetings = shakeCoords(meetings)

  return meetings
}

func sortByCoords(meetings []models.ExtendedMeeting) []models.ExtendedMeeting {
  if len(meetings) < 2 {
    return meetings
  }

  var (
    middle = getVectorLength(meetings[0])
    left, mid, right []models.ExtendedMeeting
  )

  for _, meeting := range meetings {
    vectorLength := getVectorLength(meeting)
    if vectorLength < middle {
      left = append(left, meeting)
    } else if vectorLength == middle {
      mid = append(mid, meeting)
    } else {
      right = append(right, meeting)
    }
  }

  left = sortByCoords(left)
  right = sortByCoords(right)

  return append(left, append(mid, right...)...)
}

func getVectorLength(meeting models.ExtendedMeeting) float64 {
  return math.Sqrt(
    math.Pow(float64(meeting.PublicPlace.Latitude), 2) + math.Pow(float64(meeting.PublicPlace.Longitude), 2))
}

func shakeCoords(meetings []models.ExtendedMeeting) []models.ExtendedMeeting {
  meetingsCount := len(meetings)
  for _, border := range getMeetingsBorders(meetingsCount) {
    meetingsBlock := meetings[border.left:border.right]
    meetingsBlock = setCoordsToCenterMeetingsBlock(meetingsBlock)

    for i := border.left; i < border.right; i++ {
      meetings[i] = meetingsBlock[i-border.left]
    }
  }

  return meetings
}

func getMeetingsBorders(meetingsCount int) []Border {
  if meetingsCount <= blockSize {
    return []Border{
      {left: 0, right: meetingsCount},
    }
  }

  var borders []Border
  for i := 0; i < meetingsCount; i += blockSize {
    var leftBorder, rightBorder int
    leftBorder = i
    if i+blockSize > meetingsCount {
      rightBorder = meetingsCount
    } else {
      rightBorder = i+blockSize
    }

    borders = append(borders, Border{
      left: leftBorder,
      right: rightBorder,
    })
  }

  return borders
}

func setCoordsToCenterMeetingsBlock(meetingsBlock []models.ExtendedMeeting) []models.ExtendedMeeting {
  meetingsCount := len(meetingsBlock)
  var (
    sumLatitude, sumLongitude, avgLatitude, avgLongitude float64
  )
  for _, meeting := range meetingsBlock {
    sumLatitude += float64(meeting.Latitude)
    sumLongitude += float64(meeting.Longitude)
  }
  avgLatitude = sumLatitude / float64(meetingsCount)
  avgLongitude = sumLongitude / float64(meetingsCount)

  for meetingIdx := range meetingsBlock {
    meetingsBlock[meetingIdx].Latitude = models.Latitude(avgLatitude)
    meetingsBlock[meetingIdx].Longitude = models.Longitude(avgLongitude)
  }

  return meetingsBlock
}
