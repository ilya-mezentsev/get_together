package messages

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"internal_errors"
	mock "mock/repositories"
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

	mock.DropTables(db)
	repository = New(db)
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
	res := t.Run()
	mock.DropTables(db)
	os.Exit(res)
}

func TestRepository_SaveSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.Save(mock.GetAllMessages()[0])

	utils.AssertNil(err, t)
}

func TestRepository_SaveChatNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.Save(mock.GetMessageWithNotExistsChatId())

	utils.AssertErrorsEqual(internal_errors.UnableToFindChatById, err, t)
}

func TestRepository_SaveUserIdNotFound(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.Save(mock.GetMessageWithNotExistsUserId())

	utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, t)
}

func TestRepository_SaveSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.Save(mock.GetAllMessages()[0])

	utils.AssertNotNil(err, t)
}

func TestRepository_GetLastSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	messagesCount := 2
	messages, err := repository.GetLastMessages(1, uint(messagesCount))

	utils.AssertNil(err, t)
	utils.AssertTrue(len(messages) <= messagesCount, t)
	for messageIndex := 1; messageIndex < len(messages); messageIndex++ {
		currentMessageSendingTime, previousMessageSendingTime :=
			messages[messageIndex].SendingTime, messages[messageIndex-1].SendingTime
		currentMessageIsBeforePrevious := currentMessageSendingTime.Before(previousMessageSendingTime)
		currentMessageSendingTimeIsEqualPrevious := currentMessageSendingTime.Equal(previousMessageSendingTime)

		utils.AssertTrue(currentMessageIsBeforePrevious || currentMessageSendingTimeIsEqualPrevious, t)
	}
}

func TestRepository_GetLastSomeError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetLastMessages(1, 1)

	utils.AssertNotNil(err, t)
}

func TestRepository_GetLastMessagesAfter(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	messagesCount := 10
	messages, err := repository.GetLastMessagesAfter(1, 1, uint(messagesCount))

	utils.AssertNil(err, t)
	utils.AssertTrue(len(messages) <= messagesCount, t)
	for _, message := range messages {
		utils.AssertTrue(message.SendingTime.After(mock.GetFirstMessageSendingTime()), t)
	}
}
