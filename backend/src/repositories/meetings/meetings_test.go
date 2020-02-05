package meetings

import (
  "fmt"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
  "internal_errors"
  mock "mock/repositories"
  mock2 "mock/services"
  "models"
  "os"
  "testing"
  "utils"
)

var (
  db *sqlx.DB
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

  mock.DropTables(db)
  repository = New(db)
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
  res := t.Run()
  mock.DropTables(db)
  os.Exit(res)
}

func TestRepository_GetFullMeetingInfoSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  info, err := repository.GetFullMeetingInfo(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.FirstLabeledPlace == info.LabeledPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.FirstLabeledPlace, Got: info.LabeledPlace}))
    t.Fail()
  })
}

func TestRepository_GetFullMeetingInfoMeetingNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  _, err := repository.GetFullMeetingInfo(11)
  utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_GetFullMeetingInfoNoTableError(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetFullMeetingInfo(11)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}

func TestRepository_GetPublicMeetingsSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  meetings, err := repository.GetPublicMeetings()
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  for idx, meeting := range meetings {
    utils.Assert(mock.PublicPlaces[idx] == meeting.PublicPlace, func() {
      t.Log(
        utils.GetExpectationString(
          utils.Expectation{Expected: mock.PublicPlaces[idx], Got: meeting.PublicPlace}))
      t.Fail()
    })
  }
}

func TestRepository_GetPublicMeetingsNoTableError(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetPublicMeetings()
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}

func TestRepository_GetExtendedMeetingsSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  meetings, err := repository.GetExtendedMeetings(models.UserMeetingStatusesData{
    UserId: 1,
    Invited: "invited",
    NotInvited: "not-invited",
  })
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  for idx, meeting := range meetings {
    utils.Assert(mock.FirstUserStatuses[idx] == meeting.CurrentUserStatus, func() {
      t.Log(
        utils.GetExpectationString(
          utils.Expectation{Expected: mock.FirstUserStatuses[idx], Got: meeting.CurrentUserStatus}))
      t.Fail()
    })
  }
}

func TestRepository_GetExtendedMeetingsNoTableError(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetExtendedMeetings(models.UserMeetingStatusesData{
    UserId: 1,
    Invited: "invited",
    NotInvited: "not-invited",
  })
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}

func TestRepository_CreateMeetingSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.CreateMeeting(1, mock2.NewMeetingSettings)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  meeting, _ := repository.GetFullMeetingInfo(3)
  utils.Assert(mock2.NewMeetingSettings.LabeledPlace == meeting.LabeledPlace, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock2.NewMeetingSettings.LabeledPlace, Got: meeting.LabeledPlace}))
    t.Fail()
  })
}

func TestRepository_CreateMeetingAdminNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.CreateMeeting(11, mock2.NewMeetingSettings)
  utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_CreateMeetingNoTableError(t *testing.T) {
  mock.DropTables(db)

  err := repository.CreateMeeting(11, mock2.NewMeetingSettings)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}

func TestRepository_DeleteMeetingSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.DeleteMeeting(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  _, err = repository.GetFullMeetingInfo(1)
  utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_DeleteMeetingMeetingNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.DeleteMeeting(11)
  utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_DeleteMeetingNoTableError(t *testing.T) {
  mock.DropTables(db)

  err := repository.DeleteMeeting(11)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}

func TestRepository_UpdatedSettingsSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.UpdatedSettings(2, mock2.NewMeetingSettings)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  meeting, _ := repository.GetFullMeetingInfo(2)
  utils.Assert(mock2.NewMeetingSettings.Title == meeting.Title, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock2.NewMeetingSettings, Got: meeting}))
    t.Fail()
  })
}

func TestRepository_UpdatedSettingsMeetingNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.UpdatedSettings(11, mock2.NewMeetingSettings)
  utils.AssertErrorsEqual(internal_errors.UnableToFindByMeetingId, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_UpdatedSettingsNoTableError(t *testing.T) {
  mock.DropTables(db)

  err := repository.UpdatedSettings(11, mock2.NewMeetingSettings)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Got: err}))
    t.Fail()
  })
}
