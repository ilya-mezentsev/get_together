package meetings

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"internal_errors"
	"models"
)

const (
	adminIdNotExistsMessage        = `pq: insert or update on table "meetings" violates foreign key constraint "meetings_admin_id_fkey"`
	deleteMeetingIdNotFoundMessage = `sql: no rows in result set`

	FullMeetingInfoQuery = `
  SELECT m.id, m.admin_id, m.created_at, mp.label, mp.latitude, mp.longitude,
  ms.title, ms.description, ms.tags, ms.date_time, ms.request_description_required,
  ms.duration, ms.min_age, ms.gender, ms.max_users
  FROM meetings m
  JOIN meetings_places mp ON m.id = mp.meeting_id
  JOIN meetings_settings ms ON m.id = ms.meeting_id
  WHERE m.id = $1`

	PublicMeetingInfoQuery = `
  SELECT m.id, m.admin_id, m.created_at, mp.latitude, mp.longitude, ms.title, ms.description, ms.tags
  FROM meetings m
  JOIN meetings_places mp ON m.id = mp.meeting_id
  JOIN meetings_settings ms ON m.id = ms.meeting_id`

	ExtendedMeetingInfoQuery = `
  SELECT m.id, m.admin_id, m.created_at, mp.latitude, mp.longitude,
  ms.title, ms.description, ms.tags, ms.date_time, ms.request_description_required,
  CASE
    WHEN :user_id = ANY(m.user_ids) THEN :invited
    ELSE :not_invited
  END as current_user_status
  FROM meetings m
  JOIN meetings_places mp ON m.id = mp.meeting_id
  JOIN meetings_settings ms ON m.id = ms.meeting_id`

	AddMeetingQuery = `
  INSERT INTO meetings(admin_id, user_ids) VALUES($1, $2) RETURNING id`
	AddMeetingSettingsQuery = `
  INSERT INTO meetings_settings(
  meeting_id, title, max_users, tags, date_time, description, duration, min_age, gender, request_description_required)
  VALUES(
  :meeting_id, :title, :max_users, :tags, :date_time, :description, :duration, :min_age, :gender, :request_description_required)`
	AddMeetingPlaceQuery = `
  INSERT INTO meetings_places(meeting_id, label, latitude, longitude) VALUES(:meeting_id, :label, :latitude, :longitude)`

	DeleteMeetingQuery         = `DELETE FROM meetings WHERE id = $1 RETURNING id`
	UpdateMeetingSettingsQuery = `
  UPDATE meetings_settings
  SET title = :title, max_users = :max_users, tags = :tags, date_time = :date_time, description = :description,
  duration = :duration, min_age = :min_age, gender = :gender, request_description_required = :request_description_required
  WHERE meeting_id = :meeting_id`

	AddUserIdToMeetingQuery  = `UPDATE meetings SET user_ids = array_append(user_ids, :user_id) WHERE id = :meeting_id`
	MeetingHasUserQuery      = `SELECT 1 FROM meetings WHERE id = :meeting_id AND :user_id = ANY(user_ids)`
	KickUserFromMeetingQuery = `UPDATE meetings SET user_ids = array_remove(user_ids, :user_id) WHERE id = :meeting_id`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) GetFullMeetingInfo(meetingId uint) (models.PrivateMeeting, error) {
	var info = models.PrivateMeeting{
		LabeledPlace: &models.LabeledPlace{},
	}
	rows, err := r.db.Query(FullMeetingInfoQuery, meetingId)
	if err != nil {
		return info, err
	}

	if rows.Next() {
		err = rows.Scan(
			&info.Id, &info.AdminId, &info.CreatedAt, &info.Label, &info.Latitude, &info.Longitude,
			&info.Title, &info.Description, pq.Array(&info.Tags), &info.DateTime, &info.RequestDescriptionRequired,
			&info.Duration, &info.MinAge, &info.Gender, &info.MaxUsers)
		if err != nil {
			return info, err
		}
	} else {
		return info, internal_errors.UnableToFindMeetingById
	}

	return info, nil
}

func (r Repository) GetPublicMeetings() ([]models.PublicMeeting, error) {
	var meetings []models.PublicMeeting
	rows, err := r.db.Query(PublicMeetingInfoQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var meeting = models.PublicMeeting{
			PublicPlace: &models.PublicPlace{},
		}
		err = rows.Scan(
			&meeting.Id, &meeting.AdminId, &meeting.CreatedAt, &meeting.Latitude, &meeting.Longitude,
			&meeting.Title, &meeting.Description, pq.Array(&meeting.Tags))
		if err != nil {
			return nil, err
		}

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

func (r Repository) GetExtendedMeetings(
	userStatusesData models.UserMeetingStatusesData) ([]models.ExtendedMeeting, error) {
	var meetings []models.ExtendedMeeting
	rows, err := r.db.NamedQuery(ExtendedMeetingInfoQuery, userStatusesData)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var meeting = models.ExtendedMeeting{
			PublicPlace: &models.PublicPlace{},
		}
		err = rows.Scan(
			&meeting.Id, &meeting.AdminId, &meeting.CreatedAt, &meeting.Latitude, &meeting.Longitude,
			&meeting.Title, &meeting.Description, pq.Array(&meeting.Tags), &meeting.DateTime,
			&meeting.RequestDescriptionRequired, &meeting.CurrentUserStatus)
		if err != nil {
			return nil, err
		}

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

func (r Repository) CreateMeeting(adminId uint, settings models.AllSettings) error {
	addedMeetingId, err := r.addMeeting(adminId)
	if err != nil {
		return err
	}

	return r.addMeetingSettings(addedMeetingId, settings)
}

func (r Repository) addMeeting(adminId uint) (uint, error) {
	var addedMeetingId uint
	err := r.db.QueryRow(AddMeetingQuery, adminId, pq.Array([]uint{adminId})).Scan(&addedMeetingId)

	switch {
	case err == nil:
		return addedMeetingId, nil
	case err.Error() == adminIdNotExistsMessage:
		return 0, internal_errors.UnableToFindUserById
	default:
		return 0, err
	}
}

func (r Repository) addMeetingSettings(meetingId uint, settings models.AllSettings) error {
	meeting := r.meetingSettingsToMap(meetingId, settings)
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(AddMeetingSettingsQuery, meeting)
	if err != nil {
		return err
	}
	_, err = tx.NamedExec(AddMeetingPlaceQuery, meeting)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r Repository) meetingSettingsToMap(meetingId uint, settings models.AllSettings) map[string]interface{} {
	return map[string]interface{}{
		"meeting_id":                   meetingId,
		"title":                        settings.Title,
		"max_users":                    settings.MaxUsers,
		"tags":                         pq.Array(settings.Tags),
		"date_time":                    settings.DateTime,
		"description":                  settings.Description,
		"duration":                     settings.Duration,
		"min_age":                      settings.MinAge,
		"gender":                       settings.Gender,
		"request_description_required": settings.RequestDescriptionRequired,
		"label":                        settings.Label,
		"latitude":                     settings.Latitude,
		"longitude":                    settings.Longitude,
	}
}

func (r Repository) DeleteMeeting(meetingId uint) error {
	var deletedId uint
	err := r.db.QueryRow(DeleteMeetingQuery, meetingId).Scan(&deletedId)

	switch {
	case err == nil:
		return nil
	case err.Error() == deleteMeetingIdNotFoundMessage:
		return internal_errors.UnableToFindMeetingById
	default:
		return err
	}
}

func (r Repository) UpdateSettings(meetingId uint, settings models.AllSettings) error {
	meeting := r.meetingSettingsToMap(meetingId, settings)
	res, err := r.db.NamedExec(UpdateMeetingSettingsQuery, meeting)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return internal_errors.UnableToFindMeetingById
	}

	return nil
}

func (r Repository) AddUserToMeeting(meetingId, userId uint) error {
	userInMeeting, err := r.meetingHasUser(meetingId, userId)
	if err != nil {
		return err
	}
	if userInMeeting {
		return internal_errors.UserAlreadyInMeeting
	}

	return r.updateMeetingUserIds(AddUserIdToMeetingQuery, meetingId, userId)
}

func (r Repository) meetingHasUser(meetingId, userId uint) (bool, error) {
	rows, err := r.db.NamedQuery(MeetingHasUserQuery, r.getNamedArguments(meetingId, userId))
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func (r Repository) getNamedArguments(meetingId, userId uint) map[string]interface{} {
	return map[string]interface{}{
		"user_id":    userId,
		"meeting_id": meetingId,
	}
}

func (r Repository) updateMeetingUserIds(query string, meetingId, userId uint) error {
	res, err := r.db.NamedExec(query, r.getNamedArguments(meetingId, userId))
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return internal_errors.UnableToFindMeetingById
	}

	return nil
}

func (r Repository) KickUserFromMeeting(meetingId, userId uint) error {
	userInMeeting, err := r.meetingHasUser(meetingId, userId)
	if err != nil {
		return err
	}
	if !userInMeeting {
		return internal_errors.UserNotInMeeting
	}

	return r.updateMeetingUserIds(KickUserFromMeetingQuery, meetingId, userId)
}
