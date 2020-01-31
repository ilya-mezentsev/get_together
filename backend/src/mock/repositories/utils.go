package repositories

import "github.com/jmoiron/sqlx"

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
    tag VARCHAR(100) NOT NULL
  );

  CREATE TABLE IF NOT EXISTS meetings(
    id SERIAL PRIMARY KEY,
    admin_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_ids INTEGER[] DEFAULT NULL,
    status MEETING_STATUS DEFAULT 'pending',
    archived_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE IF NOT EXISTS meetings_settings(
    id SERIAL PRIMARY KEY,
    meeting_id INTEGER NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    max_users INTEGER DEFAULT NULL,
    tags VARCHAR(100)[] DEFAULT NULL,
    date_time TIMESTAMP NOT NULL,
    description TEXT DEFAULT '',
    duration INTEGER DEFAULT NULL,
    min_age INTEGER DEFAULT NULL,
    gender GENDER DEFAULT NULL,
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
)

func dropTables(db *sqlx.DB) {
  _, err := db.Exec(DropTablesQuery)
  if err != nil {
    panic(err)
  }
}

func initTables(db *sqlx.DB) {
  _, err := db.Exec(CreateTablesQuery)
  if err != nil {
    panic(err)
  }
}
