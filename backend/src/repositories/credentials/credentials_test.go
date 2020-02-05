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
}

// we need this function to avoiding DB errors due parallel queries
func TestMain(t *testing.M) {
  res := t.Run()
  mock.DropTables(db)
  os.Exit(res)
}

func TestRepository_GetUserIdByCredentialsSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

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
  mock.InitTables(db)
  defer mock.DropTables(db)

  _, err := repository.GetUserIdByCredentials(mock.NotExistsUser)

  utils.AssertErrorsEqual(internal_errors.UnableToLoginUserNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_GetUserIdByCredentialsErrorNoTable(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetUserIdByCredentials(mock.Users[0])
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "some error", Got: err}))
    t.Fail()
  })
}

func TestRepository_CreateUserSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

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
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.CreateUser(mock.Users[0])
  utils.AssertErrorsEqual(internal_errors.UnableToRegisterUserEmailExists, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_UpdateUserPasswordSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  user := mock.Users[0]
  user.Password = "new_pass"
  err := repository.UpdateUserPassword(user)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })

  id, _ := repository.GetUserIdByCredentials(user)
  utils.Assert(1 == id, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: 1, Got: id}))
    t.Fail()
  })
}

func TestRepository_UpdateUserPasswordUserNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  err := repository.UpdateUserPassword(mock.NotExistsUser)
  utils.AssertErrorsEqual(internal_errors.UnableToChangePasswordUserNotFound, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_UpdateUserPasswordErrorNoTable(t *testing.T) {
  mock.DropTables(db)

  err := repository.UpdateUserPassword(mock.Users[0])
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "some error", Got: err}))
    t.Fail()
  })
}

func TestRepository_GetUserEmailSuccess(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  email, err := repository.GetUserEmail(1)
  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.Users[0].Email == email, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.Users[0].Email, Got: email}))
    t.Fail()
  })
}

func TestRepository_GetUserEmailUserNotFoundError(t *testing.T) {
  mock.InitTables(db)
  defer mock.DropTables(db)

  _, err := repository.GetUserEmail(11)
  utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
}

func TestRepository_GetUserEmailErrorNoTable(t *testing.T) {
  mock.DropTables(db)

  _, err := repository.GetUserEmail(1)
  utils.Assert(nil != err, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: "some error", Got: err}))
    t.Fail()
  })
}
