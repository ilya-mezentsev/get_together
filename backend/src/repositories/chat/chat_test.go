package chat

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"internal_errors"
	mock "mock/repositories"
	"models"
	"os"
	"plugins/config"
	"testing"
	"utils"
)

var (
	db         *sqlx.DB
	repository Repository
)

func init() {
	utils.SkipInShortMode()

	var err error
	db, err = config.GetConfiguredConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repository = New(db)
}

func getChat(meetingId uint) models.Chat {
	var chat models.Chat
	err := db.Get(
		&chat,
		`SELECT id, type, status, created_at FROM chats WHERE meeting_id = $1 AND type = 'meeting'`,
		meetingId,
	)
	if err != nil {
		panic(err)
	}

	return chat
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
	res := t.Run()
	mock.DropTables(db)
	os.Exit(res)
}

func TestRepository_GetMeetingChatSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	chat, err := repository.GetMeetingChat(1)
	firstChat := mock.MeetingChats[0]

	utils.AssertNil(err, t)
	utils.AssertEqual(firstChat["type"].(string), chat.Type, t)
	utils.AssertNotEqual(mock.ArchivedStatus, chat.Status, t)
}

func TestRepository_GetMeetingChatNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetMeetingChat(mock.GetNotExistsMeetingId())
	utils.AssertErrorsEqual(internal_errors.UnableToFindChatByMeetingId, err, t)
}

func TestRepository_GetMeetingChatSomeError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetMeetingChat(1)
	utils.AssertNotNil(err, t)
}

func TestRepository_GetUserChatsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	chats, err := repository.GetUserChats(1)
	firstChat := mock.MeetingChats[0]

	utils.AssertNil(err, t)
	utils.AssertEqual(firstChat["type"].(string), chats[0].Type, t)
	for _, chat := range chats {
		utils.AssertNotEqual(mock.ArchivedStatus, chat.Status, t)
	}
}

func TestRepository_GetUserChatsSomeError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetUserChats(1)
	utils.AssertNotNil(err, t)
}

func TestRepository_CreateChatSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateChat(mock.MeetingIdWithoutMeetingChat, mock.MeetingType)
	chat, _ := repository.GetMeetingChat(mock.MeetingIdWithoutMeetingChat)

	utils.AssertNil(err, t)
	utils.AssertEqual(uint(len(mock.MeetingChats)+1), chat.Id, t)
}

func TestRepository_CreateChatMeetingAlreadyHasMeetingChat(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateChat(1, mock.MeetingType)

	utils.AssertErrorsEqual(internal_errors.MeetingChatAlreadyExists, err, t)
}

func TestRepository_CreateChatSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.CreateChat(1, mock.MeetingType)

	utils.AssertNotNil(err, t)
}

func TestRepository_SetChatStatusSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.SetChatStatus(1, mock.ArchivedStatus)
	chat := getChat(1)

	utils.AssertNil(err, t)
	utils.AssertEqual(mock.ArchivedStatus, chat.Status, t)
}

func TestRepository_SetChatStatusChatNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.SetChatStatus(mock.NotExistsChatId, mock.ArchivedStatus)

	utils.AssertErrorsEqual(internal_errors.UnableToFindChatById, err, t)
}

func TestRepository_SetChatStatusSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.SetChatStatus(1, mock.ArchivedStatus)

	utils.AssertNotNil(err, t)
}
