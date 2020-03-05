package services

import (
  "internal_errors"
  "models"
  "time"
)

type MeetingsRepositoryMock struct {
  Meetings map[uint]models.PrivateMeeting
}

var (
  meetings = map[uint]models.PrivateMeeting{
    1: {
      DefaultMeeting: models.DefaultMeeting{
        ID: 1,
        AdminId: 1,
        CreatedAt: time.Unix(0, 0),
      },
      LabeledPlace: models.LabeledPlace{
        Label: "221b baker street",
        PublicPlace: models.PublicPlace{
          Latitude: 51.5207,
          Longitude: -0.1550,
        },
      },
      AllSettings: models.AllSettings{
        ExtendedSettings: models.ExtendedSettings{
          PublicSettings: models.PublicSettings{
            Title: "Deduction party",
            Description: "Fuck you, Moriarty",
            Tags: []string{"logic", "smoking pipe"},
          },
          MeetingParameters: models.MeetingParameters{
            DateTime: time.Unix(0, 1),
            RequestDescriptionRequired: false,
          },
        },
        MeetingLimitations: models.MeetingLimitations{
          Duration: 2,
          MinAge: 16,
          MaxUsers: 10,
        },
      },
    },
  }
  NewMeetingSettings = models.AllSettings{
    ExtendedSettings: models.ExtendedSettings{
      PublicSettings: models.PublicSettings{
        Title: "Winx top!",
        Description: "Who likes winx come!",
        Tags: []string{"winx", "my_love"},
      },
      MeetingParameters: models.MeetingParameters{
        DateTime: time.Unix(0, 0),
        RequestDescriptionRequired: false,
      },
    },
    LabeledPlace: models.LabeledPlace{
      Label: "221b baker street",
      PublicPlace: models.PublicPlace{
        Latitude: 51.5207,
        Longitude: -0.1550,
      },
    },
    MeetingLimitations: models.MeetingLimitations{
      Duration: 2,
      MinAge: 12,
      Gender: "female",
      MaxUsers: 10,
    },
  }
  MeetingsMockRepository = MeetingsRepositoryMock{Meetings: meetings}
  BadMeetingId uint = 0
  NotExistsMeetingId uint = 11
  NotExistsUserId uint = 11
  TestMeetingTrue := models.MeetingLimitation{5, 16, "male", 4}
  TestMeetingFalse := models.MeetingLimitation{0, 18, "female", 8}
)

func (m *MeetingsRepositoryMock) ResetState() {
  m.Meetings = meetings
}

func (m *MeetingsRepositoryMock) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
  if meetingId == BadMeetingId {
    return models.PrivateMeeting{}, someInternalError
  }

  meetingInfo, found := m.Meetings[meetingId]
  if !found {
    return models.PrivateMeeting{}, internal_errors.UnableToFindByMeetingId
  }

  return meetingInfo, nil
}

func (m *MeetingsRepositoryMock) GetPublicMeetings() ([]models.PublicMeeting, error) {
  if m.Meetings == nil {
    return nil, someInternalError
  }

  var meetings []models.PublicMeeting
  for _, m := range m.Meetings {
    meetings = append(meetings, models.PublicMeeting{
      DefaultMeeting: m.DefaultMeeting,
      PublicSettings: m.PublicSettings,
      PublicPlace: &m.PublicPlace,
    })
  }

  return meetings, nil
}

func (m *MeetingsRepositoryMock) GetExtendedMeetings(
  data models.UserMeetingStatusesData) ([]models.ExtendedMeeting, error) {
  userId := data.UserId
  if userId == BadUserId {
    return nil, someInternalError
  } else if userId == NotExistsUserId {
    return nil, internal_errors.UnableToFindUserById
  }

  var meetings []models.ExtendedMeeting
  for _, m := range m.Meetings {
    meetings = append(meetings, models.ExtendedMeeting{
      DefaultMeeting: m.DefaultMeeting,
      ExtendedSettings: m.ExtendedSettings,
      PublicPlace: &m.PublicPlace,
      CurrentUserStatus: "",
    })
  }

  return meetings, nil
}

func (m *MeetingsRepositoryMock) CreateMeeting(adminId uint, settings models.AllSettings) error {
  if adminId == BadUserId {
    return someInternalError
  } else if adminId == NotExistsUserId {
    return internal_errors.UnableToFindUserById
  }

  m.Meetings[2] = models.PrivateMeeting{
    DefaultMeeting: models.DefaultMeeting{
      ID: 2,
      AdminId: adminId,
      CreatedAt: time.Unix(0, 1),
    },
    AllSettings: settings,
  }
  return nil
}

func (m *MeetingsRepositoryMock) DeleteMeeting(meetingId uint) error {
  if meetingId == BadMeetingId {
    return someInternalError
  } else if meetingId == NotExistsMeetingId {
    return internal_errors.UnableToFindByMeetingId
  }

  delete(m.Meetings, meetingId)
  return nil
}

func (m *MeetingsRepositoryMock) UpdatedSettings(meetingId uint, settings models.AllSettings) error {
  if meetingId == BadMeetingId {
    return someInternalError
  } else if meetingId == NotExistsMeetingId {
    return internal_errors.UnableToFindByMeetingId
  }

  meeting := m.Meetings[meetingId]
  m.Meetings[meetingId] = models.PrivateMeeting{
    DefaultMeeting: meeting.DefaultMeeting,
    LabeledPlace: meeting.LabeledPlace,
    AllSettings: settings,
  }
  return nil
}
