package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"utils"
)

const (
	DropTablesQuery = `
  DROP TABLE IF EXISTS users CASCADE;
	DROP TABLE IF EXISTS users_credentials;
  DROP TABLE IF EXISTS users_info;
  DROP TABLE IF EXISTS users_rating;
  DROP TABLE IF EXISTS meetings CASCADE;
  DROP TABLE IF EXISTS meetings_settings;
  DROP TABLE IF EXISTS meetings_places;
  DROP TABLE IF EXISTS chats CASCADE;
  DROP TABLE IF EXISTS messages;
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
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS users_credentials(
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
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

	CREATE TABLE IF NOT EXISTS messages(
		id SERIAL PRIMARY KEY,
		chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
		sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		text TEXT NOT NULL,
		sending_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	CreateUserQuery            = `INSERT INTO users DEFAULT VALUES;`
	CreateUserCredentialsQuery = `
	INSERT INTO users_credentials(user_id, email, password)
	VALUES(:user_id, :email, :password)`
	CreateUserInfoQuery = `
  INSERT INTO users_info(user_id, name, nickname, age, gender)
  VALUES(:user_id, :name, :nickname, :age, :gender);`
	CreateUserRatingQuery      = `INSERT INTO users_rating(user_id, tag, value) VALUES(:user_id, :tag, :value);`
	CreateMeetingQuery         = `INSERT INTO meetings(admin_id, user_ids) VALUES(:admin_id, :user_ids);`
	CreateMeetingSettingsQuery = `
  INSERT INTO meetings_settings(meeting_id, title, date_time, tags, duration, max_users, min_age, gender)
  VALUES(:meeting_id, :title, :date_time, :tags, :duration, :max_users, :min_age, :gender);`
	CreateMeetingPlaceQuery = `
  INSERT INTO meetings_places(meeting_id, label, latitude, longitude)
  VALUES(:meeting_id, :label, :latitude, :longitude);`
	CreateChatQuery    = `INSERT INTO chats(meeting_id, type) VALUES(:meeting_id, :type)`
	CreateMessageQuery = `
	INSERT INTO messages(chat_id, sender_id, text) VALUES(:chat_id, :sender_id, :text)`
)

var (
	TestingPassword = "mYStRoNg*PwD12"
	Users           = []map[string]interface{}{
		{}, {}, {}, {}, // using default values in query, so no need data here
	}
	UsersCredentials = []map[string]interface{}{
		{"user_id": 1, "email": "mail@ya.ru"}, {"user_id": 2, "email": "me@gmail.com"},
		{"user_id": 3, "email": "hello.world@mail.ru"}, {"user_id": 4, "email": "world@hello.ru"},
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
		{"meeting_id": 1, "admin_id": 1, "user_ids": []uint{1}},
		{"meeting_id": 2, "admin_id": 2, "user_ids": []uint{2, 4}},
		{"meeting_id": 3, "admin_id": 3, "user_ids": []uint{3}},
	}
	MeetingsSettings = []map[string]interface{}{
		{
			"meeting_id":                   1,
			"title":                        "hello_world",
			"date_time":                    "2020-03-02T14:00:00",
			"tags":                         pq.Array([]string{"tag1", "tag2"}),
			"duration":                     4,
			"max_users":                    10,
			"min_age":                      16,
			"gender":                       "male",
			"request_description_required": true,
		},
		{
			"meeting_id":                   2,
			"title":                        "hello_world",
			"date_time":                    "2020-03-02T16:00:00",
			"tags":                         pq.Array([]string{"tag3"}),
			"duration":                     4,
			"max_users":                    5,
			"min_age":                      18,
			"gender":                       "female",
			"request_description_required": false,
		},
		{
			"meeting_id":                   3,
			"title":                        "hello_world",
			"date_time":                    "2020-03-02T20:00:00",
			"tags":                         pq.Array([]string{"tag1"}),
			"duration":                     4,
			"max_users":                    6,
			"min_age":                      12,
			"gender":                       "",
			"request_description_required": true,
		},
	}
	MeetingsPlaces = []map[string]interface{}{
		{"meeting_id": 1, "label": "221b baker street", "latitude": 51.5207, "longitude": -0.1550},
		{"meeting_id": 2, "label": "hello-world", "latitude": 0.0, "longitude": 0.0},
		{"meeting_id": 3, "label": "221b baker street", "latitude": 51.5207, "longitude": -0.1550},
	}
	MeetingChats = []map[string]interface{}{
		{"meeting_id": 1, "type": "meeting"}, {"meeting_id": 1, "type": "meeting_request"},
		{"meeting_id": 2, "type": "meeting"}, {"meeting_id": 2, "type": "meeting_request"},
		{"meeting_id": 3, "type": "meeting_request"},
	}
	ChatsMessages = []map[string]interface{}{
		{"chat_id": 1, "sender_id": 1, "text": "hello world 1"},
		{"chat_id": 5, "sender_id": 1, "text": "hello world 3"},
		{"chat_id": 1, "sender_id": 2, "text": "hello world 4"},
		{"chat_id": 5, "sender_id": 2, "text": "hello world 6"},
		{"chat_id": 1, "sender_id": 3, "text": "hello world 7"},
		{"chat_id": 5, "sender_id": 3, "text": "hello world 9"},
	}
	QueryToSubData = map[string][]map[string]interface{}{
		CreateUserCredentialsQuery: UsersCredentials,
		CreateUserInfoQuery:        UsersInfo,
		CreateUserRatingQuery:      UsersRating,
		CreateMeetingSettingsQuery: MeetingsSettings,
		CreateMeetingPlaceQuery:    MeetingsPlaces,
		CreateMessageQuery:         ChatsMessages,
	}
)

func init() {
	for idx, c := range UsersCredentials {
		UsersCredentials[idx]["password"] = utils.GetHash(c["email"].(string) + TestingPassword)
	}
}

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

	var meetings []map[string]interface{}
	for _, m := range Meetings {
		meetings = append(meetings, map[string]interface{}{
			"meeting_id": m["meeting_id"], "admin_id": m["admin_id"], "user_ids": pq.Array(m["user_ids"]),
		})
	}

	addDataFromSource(tx, CreateUserQuery, Users)
	addDataFromSource(tx, CreateMeetingQuery, meetings)
	addDataFromSource(tx, CreateChatQuery, MeetingChats)
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
