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
  CreateDataQuery = `
  INSERT INTO users(email, password) VALUES('mail@ya.ru', '3dac4de4c9d5af7382da4c63f5555f2b');
  INSERT INTO users(email, password) VALUES('me@gmail.com', '0c120226ef10689396a6eabbf733e54b');

  INSERT INTO users_info(user_id, name, nickname, gender, age) VALUES(1, 'J. Smith', 'mather_fucker', 'male', 12);
  INSERT INTO users_info(user_id, name, nickname, age, avatar_url) VALUES(2, 'Mr. Anderson', 'LoL228', 8, 'http://123.png');

  INSERT INTO users_rating(user_id, tag, value) VALUES(1, 'tag1', 85);
  INSERT INTO users_rating(user_id, tag, value) VALUES(1, 'tag2', 75);
  INSERT INTO users_rating(user_id, tag, value) VALUES(2, 'tag1', 95);
  INSERT INTO users_rating(user_id, tag, value) VALUES(2, 'tag3', 65);

  INSERT INTO meetings(admin_id, user_ids) VALUES(1, ARRAY[1]);
  INSERT INTO meetings(admin_id, user_ids) VALUES(2, ARRAY[2]);

  INSERT INTO meetings_settings(meeting_id, title, date_time, tags)
  VALUES(1, 'hello_world', '2020-02-02T17:53:38.218Z', ARRAY['tag1', 'tag2']);
  INSERT INTO meetings_settings(meeting_id, title, date_time, tags)
  VALUES(2, 'hello_world', '2020-03-02T17:53:38.218Z', ARRAY['tag3']);

  INSERT INTO meetings_places(meeting_id, label, latitude, longitude) VALUES(1, '221b baker street', 51.5207, -0.1550);
  INSERT INTO meetings_places(meeting_id, label, latitude, longitude) VALUES(2, 'hello-world', 0.0, 0.0);
  `
)

func DropTables(db *sqlx.DB) {
  _, err := db.Exec(DropTablesQuery)
  if err != nil {
    panic(err)
  }
}

func InitTables(db *sqlx.DB) {
  _, err := db.Exec(CreateTablesQuery)
  if err != nil {
    panic(err)
  }

  _, err = db.Exec(CreateDataQuery)
  if err != nil {
    panic(err)
  }
}
