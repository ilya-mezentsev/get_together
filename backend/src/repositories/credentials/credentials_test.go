package credentials

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
  mock.InitUsers(db)
}

func TestRepository_GetUserIdByCredentialsSuccess(t *testing.T) {
  mock.InitUsers(db)
  defer mock.DropUsers(db)

  id, err := repository.GetUserIdByCredentials(mock.Users[0])

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(1 == id, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: 1, Got: id}))
    t.Fail()
  })
}

func TestRepository_GetUserIdByCredentialsErrorUserNotExists(t *testing.T) {
  mock.InitUsers(db)
  defer mock.DropUsers(db)

  _, err := repository.GetUserIdByCredentials(mock.NotExistsUser)

  utils.AssertErrorsEqual(internal_errors.UnableToLoginUserNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_GetUserIdByCredentialsErrorNoTable(t *testing.T) {
  mock.DropUsers(db)

  _, err := repository.GetUserIdByCredentials(mock.Users[0])
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "some error", Got: err}))
    t.Fail()
  })
}

func TestRepository_CreateUserSuccess(t *testing.T) {
  mock.InitUsers(db)
  defer mock.DropUsers(db)

  err := repository.CreateUser(mock.NewUser)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })

  id, _ := repository.GetUserIdByCredentials(mock.NewUser)
  utils.Assert(3 == id, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: 3, Got: id}))
    t.Fail()
  })
}

func TestRepository_CreateUserEmailExistsError(t *testing.T) {
  mock.InitUsers(db)
  defer mock.DropUsers(db)

  err := repository.CreateUser(mock.Users[0])
  utils.AssertErrorsEqual(internal_errors.UnableToRegisterUserEmailExists, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_CreateUserIncorrectPasswordError(t *testing.T) {
  mock.InitUsers(db)
  defer mock.DropUsers(db)

  err := repository.CreateUser(mock.IncorrectPasswordUser)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "some error", Got: err}))
    t.Fail()
  })
}
