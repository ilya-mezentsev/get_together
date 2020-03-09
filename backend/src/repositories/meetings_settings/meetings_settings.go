package meetings_settings

import (
  "github.com/jmoiron/sqlx"
  "github.com/lib/pq"
  "internal_errors"
  "models"
)

const (
  GetMeetingSettingsQuery = `
  SELECT max_users, tags, date_time, duration, min_age, gender, request_description_required,
  CARDINALITY(m.user_ids) as users_count
  FROM meetings_settings ms
  JOIN meetings m ON m.id = ms.meeting_id
  WHERE ms.meeting_id = $1
  `
  GetNearMeetingsQuery = `
  SELECT ms.date_time, ms.duration
  FROM meetings m
  JOIN meetings_settings ms ON AGE(ms.date_time, (
    SELECT date_time FROM meetings_settings WHERE meeting_id = :meeting_id
  )) < interval '1 day' AND m.id = ms.meeting_id
  WHERE :user_id = ANY(m.user_ids)
  `
)

type Repository struct {
  db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
  return Repository{db}
}

func (r Repository) GetMeetingSettings(meetingId uint) (models.ParticipationMeetingSettings, error) {
  var settings models.ParticipationMeetingSettings
  rows, err := r.db.Query(GetMeetingSettingsQuery, meetingId)
  if err != nil {
    return models.ParticipationMeetingSettings{}, err
  }

  if rows.Next() {
    err = rows.Scan(
      &settings.MaxUsers, pq.Array(&settings.Tags), &settings.DateTime, &settings.Duration,
      &settings.MinAge, &settings.Gender, &settings.RequestDescriptionRequired, &settings.UsersCount)
  } else {
    return models.ParticipationMeetingSettings{}, internal_errors.UnableToFindByMeetingId
  }

  return settings, err
}

func (r Repository) GetNearMeetings(data models.UserTimeCheckData) ([]models.TimeMeetingParameters, error) {
  var meetings []models.TimeMeetingParameters
  rows, err := r.db.NamedQuery(GetNearMeetingsQuery, data)
  if err != nil {
    return nil, err
  }

  for rows.Next() {
    var meeting models.TimeMeetingParameters
    err = rows.StructScan(&meeting)

    meetings = append(meetings, meeting)
  }

  return meetings, err
}
