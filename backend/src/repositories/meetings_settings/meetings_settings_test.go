package meetings_settings

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"internal_errors"
	mock "mock/repositories"
	mock2 "mock/services"
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

func TestRepository_GetMeetingSettingsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	settings, err := repository.GetMeetingSettings(1)
	utils.AssertNil(err, t)
	utils.AssertEqual(uint(1), settings.UsersCount, t)
}

func TestRepository_GetMeetingSettingsMeetingNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetMeetingSettings(mock2.NotExistsMeetingId)
	utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, t)
}

func TestRepository_GetMeetingSettingsTableNotFoundError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetMeetingSettings(1)
	utils.AssertNotNil(err, t)
}

func TestRepository_GetBeforeMeetingsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	meetings, err := repository.GetNearMeetings(mock.TimeCheckData)
	utils.AssertNil(err, t)
	utils.AssertTrue(meetings[0].DateTime.Equal(mock.ExpectedDate), t)
}

func TestRepository_GetNearMeetingsTableNotFoundError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetNearMeetings(mock.TimeCheckData)
	utils.AssertNotNil(err, t)
}
