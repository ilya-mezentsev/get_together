package services

import (
  "internal_errors"
  "models"
  "time"
)

type MeetingsSettingsRepositoryMock struct {
  meetings map[uint]models.ParticipationMeetingSettings
}

var (
  meetingsSettings = map[uint]models.ParticipationMeetingSettings{
    1: {
      MeetingLimitations: models.MeetingLimitations{
        MaxUsers: 10,
        Duration: 4,
        MinAge: 16,
        Gender: "male",
      },
      MeetingParameters: models.MeetingParameters{
        DateTime: time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC),
        RequestDescriptionRequired: false,
      },
      Tags: []string{"tag1", "tag2"},
      UsersCount: 2,
    },
    2: {
      MeetingLimitations: models.MeetingLimitations{
        MaxUsers: 5,
        Duration: 4,
        MinAge: 12,
        Gender: "female",
      },
      MeetingParameters: models.MeetingParameters{
        DateTime: time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC),
        RequestDescriptionRequired: false,
      },
      Tags: []string{"tag3"},
      UsersCount: 5,
    },
    3:  {
      MeetingLimitations: models.MeetingLimitations{
        MaxUsers: 6,
        Duration: 2,
        MinAge: 18,
        Gender: "male",
      },
      MeetingParameters: models.MeetingParameters{
        DateTime: time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC),
        RequestDescriptionRequired: false,
      },
      Tags: []string{"tag3"},
      UsersCount: 1,
    },
    4: {
      MeetingLimitations: models.MeetingLimitations{
        MaxUsers: 5,
        Duration: 4,
        MinAge: 12,
        Gender: "female",
      },
      MeetingParameters: models.MeetingParameters{
        DateTime: time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC),
        RequestDescriptionRequired: false,
      },
      Tags: []string{"tag3"},
      UsersCount: 2,
    },
  }
  MeetingsSettingsRepository = MeetingsSettingsRepositoryMock{
    meetings: meetingsSettings,
  }
  MeetingIdWithParticipation uint = 3

  HasNearMeetingRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: MeetingIdWithParticipation,
  }
  BadRatingRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: 1,
  }
  MeetingFullRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: 2,
  }
  InappropriateAgeRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: 3,
  }
  WrongGenderRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: 4,
  }
  NotExistsUserIdRequest = models.ParticipationRequest{
    UserId: NotExistsUserId,
    MeetingId: 4,
  }
  NotExistsMeetingIdRequest = models.ParticipationRequest{
    UserId: 1,
    MeetingId: NotExistsMeetingId,
  }
  InternalErrorRequest1 = models.ParticipationRequest{
    UserId: BadUserId,
    MeetingId: 1,
  }
  InternalErrorRequest2 = models.ParticipationRequest{
    UserId: 1,
    MeetingId: BadMeetingId,
  }
  WrongGender = models.InappropriateInfoField{
    ErrorCode: "wrong-gender",
    Description: "actual: male, wanted: female",
  }
  InappropriateAge = models.InappropriateInfoField{
    ErrorCode: "age-less-than-min",
    Description: "actual: 16, wanted: 18",
  }
  MaxUsersCountReached = models.InappropriateInfoField{
    ErrorCode: "max-users-count-reached",
    Description: "actual: 5",
  }
)

func (m *MeetingsSettingsRepositoryMock) ResetState() {
  m.meetings = meetingsSettings
}

func (m *MeetingsSettingsRepositoryMock) GetMeetingSettings(meetingId uint) (models.ParticipationMeetingSettings, error) {
  if meetingId == BadMeetingId {
    return models.ParticipationMeetingSettings{}, someInternalError
  } else if meetingId == NotExistsMeetingId {
    return models.ParticipationMeetingSettings{}, internal_errors.UnableToFindByMeetingId
  }

  return m.meetings[meetingId], nil
}

func (m *MeetingsSettingsRepositoryMock) GetNearMeetings(data models.UserTimeCheckData) ([]models.TimeMeetingParameters, error) {
  if data.MeetingId == BadMeetingId {
    return nil, someInternalError
  } else if data.MeetingId == NotExistsMeetingId {
    return nil, internal_errors.UnableToFindByMeetingId
  } else if data.UserId == NotExistsUserId {
    return nil, internal_errors.UnableToFindUserById
  }

  var meetings []models.TimeMeetingParameters
  for meetingId, meeting := range meetingsSettings {
    if meetingId != data.MeetingId {
      meetings = append(meetings, models.TimeMeetingParameters{
        DateTime: meeting.DateTime,
        Duration: meeting.Duration,
      })
    }
  }
  return meetings, nil
}

func TagsEqual(t1, t2 []string) bool {
  if len(t1) != len(t2) {
    return false
  }

  for idx := range t1 {
    if t1[idx] != t2[idx] {
      return false
    }
  }

  return true
}
