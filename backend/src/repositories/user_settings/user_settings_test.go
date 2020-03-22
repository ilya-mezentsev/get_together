package user_settings

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"internal_errors"
	mock "mock/repositories"
	"os"
	"testing"
	"utils"
)

var (
	db         *sqlx.DB
	repository Repository
)

func init() {
	utils.SkipInShortMode()

	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		fmt.Println("CONN_STR env var is not set")
		os.Exit(1)
	}

	var err error
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repository = New(db)
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
	res := t.Run()
	mock.DropTables(db)
	os.Exit(res)
}

func TestRepository_GetUserInfoSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	userSettings, err := repository.GetUserSettings(1)
	utils.AssertNil(err, t)
	utils.AssertTrue(mock.SettingsEqual(mock.GetFirstUserSettings(), userSettings), t)
}

func TestRepository_GetUserInfoUserNoFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetUserSettings(mock.GetNotExistsUserId())
	utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, t)
}

func TestRepository_GetUserInfoErrorNoTable(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetUserSettings(mock.GetNotExistsUserId())
	utils.AssertNotNil(err, t)
}

func TestRepository_UpdateUserInfoSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.UpdateUserSettings(1, mock.TestInfo)
	utils.AssertNil(err, t)

	userSettings, _ := repository.GetUserSettings(1)
	utils.AssertEqual(mock.TestInfo, userSettings.UserSettings, t)
}

func TestRepository_UpdateUserInfoUserNoFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.UpdateUserSettings(mock.GetNotExistsUserId(), mock.TestInfo)
	utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, t)
}

func TestRepository_UpdateUserInfoErrorNoTable(t *testing.T) {
	mock.DropTables(db)

	err := repository.UpdateUserSettings(1, mock.TestInfo)
	utils.AssertNotNil(err, t)
}
