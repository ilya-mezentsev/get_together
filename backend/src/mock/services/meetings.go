package services

import (
	"github.com/lib/pq"
	"internal_errors"
	"mock/repositories"
	"models"
	"services/proxies/validation/plugins/validation"
	"time"
)

type MeetingsRepositoryMock struct {
	Meetings      map[uint]models.PrivateMeeting
	MeetingsUsers map[uint][]uint
}

var (
	NewMeetingSettings = models.AllSettings{
		ExtendedSettings: models.ExtendedSettings{
			PublicSettings: models.PublicSettings{
				Title:       "Winx top!",
				Description: "Who likes winx come!",
				Tags:        []string{"winx", "my_love"},
			},
			MeetingParameters: models.MeetingParameters{
				DateTime:                   time.Unix(0, 0),
				RequestDescriptionRequired: false,
			},
		},
		LabeledPlace: models.LabeledPlace{
			Label: "221b baker street",
			PublicPlace: models.PublicPlace{
				Latitude:  51.5207,
				Longitude: -0.1550,
			},
		},
		MeetingLimitations: models.MeetingLimitations{
			Duration: 2,
			MinAge:   12,
			Gender:   "female",
			MaxUsers: 10,
		},
	}
	MeetingsMockRepository = MeetingsRepositoryMock{
		Meetings:      allMeetings(),
		MeetingsUsers: allMeetingsUsers(),
	}
	UserIdThatNotInFirstMeeting      = repositories.UserIdThatNotInFirstMeeting
	BadMeetingId                uint = 0
	NotExistsMeetingId          uint = 11
	NotExistsUserId             uint = 11
)

func (m *MeetingsRepositoryMock) ResetState() {
	m.Meetings = allMeetings()
	m.MeetingsUsers = allMeetingsUsers()
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
			PublicPlace:    &m.PublicPlace,
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
			DefaultMeeting:    m.DefaultMeeting,
			ExtendedSettings:  m.ExtendedSettings,
			PublicPlace:       &m.PublicPlace,
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
			ID:        2,
			AdminId:   adminId,
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

func (m *MeetingsRepositoryMock) UpdateSettings(meetingId uint, settings models.AllSettings) error {
	if meetingId == BadMeetingId {
		return someInternalError
	} else if meetingId == NotExistsMeetingId {
		return internal_errors.UnableToFindByMeetingId
	}

	meeting := m.Meetings[meetingId]
	m.Meetings[meetingId] = models.PrivateMeeting{
		DefaultMeeting: meeting.DefaultMeeting,
		LabeledPlace:   meeting.LabeledPlace,
		AllSettings:    settings,
	}
	return nil
}

func (m *MeetingsRepositoryMock) AddUserToMeeting(meetingId, userId uint) error {
	if meetingId == BadMeetingId {
		return someInternalError
	}

	for id, userIds := range m.MeetingsUsers {
		if id == meetingId {
			if HasUser(userIds, userId) {
				return internal_errors.UserAlreadyInMeeting
			} else {
				m.MeetingsUsers[id] = append(userIds, userId)
				return nil
			}
		}
	}

	return internal_errors.UnableToFindByMeetingId
}

func (m *MeetingsRepositoryMock) KickUserFromMeeting(meetingId, userId uint) error {
	if meetingId == BadMeetingId {
		return someInternalError
	}

	for id, userIds := range m.MeetingsUsers {
		if id == meetingId {
			if HasUser(userIds, userId) {
				m.MeetingsUsers[id] = filterUserIds(userIds, userId)
				return nil
			} else {
				return internal_errors.UserNotInMeeting
			}
		}
	}

	return internal_errors.UnableToFindByMeetingId
}

func HasUser(userIds []uint, userId uint) bool {
	for _, id := range userIds {
		if id == userId {
			return true
		}
	}

	return false
}

func filterUserIds(userIds []uint, userId uint) []uint {
	var ids []uint
	for _, id := range userIds {
		if id != userId {
			ids = append(ids, id)
		}
	}

	return ids
}

func allMeetingsUsers() map[uint][]uint {
	var meetingsUsers = map[uint][]uint{}
	for _, m := range repositories.Meetings {
		meetingsUsers[uint(m["meeting_id"].(int))] = m["user_ids"].([]uint)
	}

	return meetingsUsers
}

func allMeetings() map[uint]models.PrivateMeeting {
	meetings := map[uint]models.PrivateMeeting{}
	for _, m := range repositories.MeetingsSettings {
		meetingId := uint(m["meeting_id"].(int))
		datetime, _ := time.Parse(validation.DateFormat, m["date_time"].(string))

		meetings[meetingId] = models.PrivateMeeting{
			DefaultMeeting: models.DefaultMeeting{
				ID:        meetingId,
				AdminId:   getAdminIdByMeetingId(meetingId),
				CreatedAt: time.Unix(0, 0),
			},
			LabeledPlace: getLabeledPlaceByMeetingId(meetingId),
			AllSettings: models.AllSettings{
				ExtendedSettings: models.ExtendedSettings{
					PublicSettings: models.PublicSettings{
						Title:       m["title"].(string),
						Description: "Fuck you, Moriarty",
						Tags:        pqStringArrayToStringArray(m["tags"].(*pq.StringArray)),
					},
					MeetingParameters: models.MeetingParameters{
						DateTime:                   datetime,
						RequestDescriptionRequired: false,
					},
				},
				MeetingLimitations: models.MeetingLimitations{
					Duration: uint(m["duration"].(int)),
					MinAge:   uint(m["min_age"].(int)),
					MaxUsers: uint(m["max_users"].(int)),
				},
			},
		}
	}

	return meetings
}

func getAdminIdByMeetingId(meetingId uint) uint {
	for id, m := range repositories.Meetings {
		if id+1 == int(meetingId) {
			return uint(m["admin_id"].(int))
		}
	}

	return 0
}

func getLabeledPlaceByMeetingId(meetingId uint) *models.LabeledPlace {
	for _, m := range repositories.MeetingsPlaces {
		if uint(m["meeting_id"].(int)) == meetingId {
			return &models.LabeledPlace{
				Label: m["label"].(string),
				PublicPlace: models.PublicPlace{
					Latitude:  models.Latitude(m["latitude"].(float64)),
					Longitude: models.Longitude(m["longitude"].(float64)),
				},
			}
		}
	}

	return nil
}
