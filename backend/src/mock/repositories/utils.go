package repositories

import (
  "github.com/jmoiron/sqlx"
  "github.com/lib/pq"
)

const (
  DropTablesQuery = `
  DROP TABLE IF EXISTS users CASCADE;
  DROP TABLE IF EXISTS users_info;
  DROP TABLE IF EXISTS users_rating;
  DROP TABLE IF EXISTS meetings CASCADE;
  DROP TABLE IF EXISTS meetings_settings;
  DROP TABLE IF EXISTS meetings_places;
  DROP TABLE IF EXISTS chats CASCADE;
  DROP TABLE IF EXISTS meetings_messages;
  DROP TABLE IF EXISTS meetings_request_messages;
  DROP TYPE IF EXISTS GENDER;
  DROP TYPE IF EXISTS MEETING_STATUS;
  DROP TYPE IF EXISTS CHAT_TYPE;
  DROP TYPE IF EXISTS CHAT_STATUS;`
  CreateTablesQuery = `
  CREATE TYPE GENDER AS ENUM('male', 'female', '');
  CREATE TYPE MEETING_STATUS AS ENUM('pending', 'archived');
  CREATE TYPE CHAT_TYPE AS ENUM('meeting', 'meeting_request');
  CREATE TYPE CHAT_STATUS AS ENUM('chatting', 'archived');

  CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(32) NOT NULL
  );

  CREATE TABLE IF NOT EXISTS users_info(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    gender GENDER DEFAULT '',
    age INTEGER DEFAULT 0,
    avatar_url VARCHAR DEFAULT ''
  );

  CREATE TABLE IF NOT EXISTS users_rating(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    value FLOAT NOT NULL,
    tag VARCHAR(100) NOT NULL,
    UNIQUE (user_id, tag)
  );

  CREATE TABLE IF NOT EXISTS meetings(
    id SERIAL PRIMARY KEY,
    admin_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_ids INTEGER[] NOT NULL,
    status MEETING_STATUS DEFAULT 'pending',
    archived_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE IF NOT EXISTS meetings_settings(
    id SERIAL PRIMARY KEY,
    meeting_id INTEGER NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    max_users INTEGER DEFAULT 0,
    tags VARCHAR(100)[] DEFAULT NULL,
    date_time TIMESTAMP NOT NULL,
    description TEXT DEFAULT '',
    duration INTEGER DEFAULT 0,
    min_age INTEGER DEFAULT 0,
    gender GENDER DEFAULT '',
    request_description_required BOOLEAN DEFAULT FALSE
  );

  CREATE TABLE IF NOT EXISTS meetings_places(
    id SERIAL PRIMARY KEY,
    meeting_id INTEGER NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    label VARCHAR(511) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
  );

  CREATE TABLE IF NOT EXISTS chats(
    id SERIAL PRIMARY KEY,
    meeting_id INTEGER NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    type CHAT_TYPE NOT NULL,
    status CHAT_STATUS DEFAULT 'chatting',
    archived_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE IF NOT EXISTS meetings_messages(
    id SERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    meeting_id INTEGER NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    sending_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE IF NOT EXISTS meetings_request_messages(
    id SERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    admin_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    sending_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );`
  CreateUserQuery = `INSERT INTO users(email, password) VALUES(:email, :password);`
  CreateUserInfoQuery = `
  INSERT INTO users_info(user_id, name, nickname, age, gender)
  VALUES(:user_id, :name, :nickname, :age, :gender);`
  CreateUserRatingQuery = `INSERT INTO users_rating(user_id, tag, value) VALUES(:user_id, :tag, :value);`
  CreateMeetingQuery = `INSERT INTO meetings(admin_id, user_ids) VALUES(:admin_id, :user_ids);`
  CreateMeetingSettingsQuery = `
  INSERT INTO meetings_settings(meeting_id, title, date_time, tags, duration)
  VALUES(:meeting_id, :title, :date_time, :tags, :duration);`
  CreateMeetingPlaceQuery = `
  INSERT INTO meetings_places(meeting_id, label, latitude, longitude)
  VALUES(:meeting_id, :label, :latitude, :longitude);`
)

var (
  UsersCredentials = []map[string]interface{}{
    {"email": "mail@ya.ru", "password": "3dac4de4c9d5af7382da4c63f5555f2b"},
    {"email": "me@gmail.com", "password": "0c120226ef10689396a6eabbf733e54b"},
    {"email": "hello.world@mail.ru", "password": "3dac4de4c9d5af7382da4c63f5555f2b"},
    {"email": "world@hello.ru", "password": "3dac4de4c9d5af7382da4c63f5555f2b"},
  }
  UsersInfo = []map[string]interface{}{
    {"user_id": 1, "name": "J. Smith", "nickname": "mather_fucker", "age": 12, "gender": "male"},
    {"user_id": 2, "name": "Mr. Anderson", "nickname": "Lol228", "age": 8, "gender": "male"},
    {"user_id": 3, "name": "Alex", "nickname": "nagibator", "age": 21, "gender": "female"},
  }
  UsersRating = []map[string]interface{}{
    {"user_id": 1, "tag": "tag1", "value": 65},
    {"user_id": 1, "tag": "tag2", "value": 55},
    {"user_id": 1, "tag": "tag3", "value": 43},
    {"user_id": 2, "tag": "tag1", "value": 40},
    {"user_id": 2, "tag": "tag2", "value": 90},
    {"user_id": 2, "tag": "tag3", "value": 55},
    {"user_id": 3, "tag": "tag1", "value": 40},
    {"user_id": 3, "tag": "tag2", "value": 90},
    {"user_id": 3, "tag": "tag3", "value": 55},
  }
  Meetings = []map[string]interface{}{
    {"admin_id": 1, "user_ids": pq.Array([]uint{1})},
    {"admin_id": 2, "user_ids": pq.Array([]uint{2, 4})},
    {"admin_id": 3, "user_ids": pq.Array([]uint{3})},
  }
  MeetingsSettings = []map[string]interface{}{
    {
      "meeting_id": 1,
      "title": "hello_world",
      "date_time": "2020-03-02T14:00:00",
      "tags": pq.Array([]string{"tag1", "tag2"}),
      "duration": 4,
    },
    {
      "meeting_id": 2,
      "title": "hello_world",
      "date_time": "2020-03-02T16:00:00",
      "tags": pq.Array([]string{"tag3"}),
      "duration": 4,
    },
    {
      "meeting_id": 3,
      "title": "hello_world",
      "date_time": "2020-03-02T20:00:00",
      "tags": pq.Array([]string{"tag1"}),
      "duration": 4,
    },
  }
  MeetingsPlaces = []map[string]interface{}{
    {"meeting_id": 1, "label": "221b baker street", "latitude": 51.5207, "longitude": -0.1550},
    {"meeting_id": 2, "label": "hello-world", "latitude": 0.0, "longitude": 0.0},
    {"meeting_id": 3, "label": "221b baker street", "latitude": 51.5207, "longitude": -0.1550},
  }
  QueryToSubData = map[string][]map[string]interface{}{
    CreateUserInfoQuery: UsersInfo,
    CreateUserRatingQuery: UsersRating,
    CreateMeetingSettingsQuery: MeetingsSettings,
    CreateMeetingPlaceQuery: MeetingsPlaces,
  }
)

func DropTables(db *sqlx.DB) {
  _, err := db.Exec(DropTablesQuery)
  if err != nil {
    panic(err)
  }
}

func InitTables(db *sqlx.DB) {
  DropTables(db)
  _, err := db.Exec(CreateTablesQuery)
  if err != nil {
    panic(err)
  }

  insertData(db)
}

func insertData(db *sqlx.DB) {
  tx := db.MustBegin()
  addDataFromSource(tx, CreateUserQuery, UsersCredentials)
  addDataFromSource(tx, CreateMeetingQuery, Meetings)
  if err := tx.Commit(); err != nil {
    panic(err)
  }

  tx = db.MustBegin()
  for query, data := range QueryToSubData {
    addDataFromSource(tx, query, data)
  }
  if err := tx.Commit(); err != nil {
    panic(err)
  }
}

func addDataFromSource(tx *sqlx.Tx, query string, src []map[string]interface{}) {
  for _, item := range src {
    _, err := tx.NamedExec(query, item)

    if err != nil {
      panic(err)
    }
  }
}
