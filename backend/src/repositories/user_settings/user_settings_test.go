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
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.SettingsEqual(mock.FirstUserSettings, userSettings), func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.FirstUserSettings, Got: userSettings}))
    t.Fail()
  })
}

func TestRepository_GetUserInfoUserNoFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  _, err := repository.GetUserSettings(11)
  utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_GetUserInfoErrorNoTable(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetUserSettings(11)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "not nil error", Got: err}))
    t.Fail()
  })
}

func TestRepository_UpdateUserInfoSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.UpdateUserSettings(1, mock.TestInfo)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })

  userSettings, _ := repository.GetUserSettings(1)
  utils.Assert(mock.TestInfo == userSettings.UserSettings, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.TestInfo, Got: userSettings.UserSettings}))
    t.Fail()
  })
}

func TestRepository_UpdateUserInfoUserNoFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.UpdateUserSettings(11, mock.TestInfo)
  utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_UpdateUserInfoErrorNoTable(t *testing.T) {
  mock.DropTables(db)

  err := repository.UpdateUserSettings(1, mock.TestInfo)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "not nil error", Got: err}))
    t.Fail()
  })
}
