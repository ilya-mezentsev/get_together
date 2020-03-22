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

func TestRepository_GetUserIdByCredentialsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	id, err := repository.GetUserIdByCredentials(mock.GetFirstUser())

	utils.AssertNil(err, t)
	utils.AssertEqual(1, int(id), t)
}

func TestRepository_GetUserIdByCredentialsErrorUserNotExists(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetUserIdByCredentials(mock.NotExistsUser)

	utils.AssertErrorsEqual(internal_errors.UnableToLoginUserNotFound, err, t)
}

func TestRepository_GetUserIdByCredentialsErrorNoTable(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetUserIdByCredentials(mock.GetFirstUser())
	utils.AssertNotNil(err, t)
}

func TestRepository_CreateUserSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateUser(mock.NewUser)
	utils.AssertNil(err, t)

	id, _ := repository.GetUserIdByCredentials(mock.NewUser)
	utils.AssertEqual(mock.GetNextUserId(), id, t)
}

func TestRepository_CreateUserEmailExistsError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.CreateUser(mock.GetFirstUser())
	utils.AssertErrorsEqual(internal_errors.UnableToRegisterUserEmailExists, err, t)
}

func TestRepository_UpdateUserPasswordSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	user := mock.GetFirstUser()
	user.Password = "new_pass"
	err := repository.UpdateUserPassword(user)
	utils.AssertNil(err, t)

	id, _ := repository.GetUserIdByCredentials(user)
	utils.AssertEqual(1, int(id), t)
}

func TestRepository_UpdateUserPasswordUserNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.UpdateUserPassword(mock.NotExistsUser)
	utils.AssertErrorsEqual(internal_errors.UnableToChangePasswordUserNotFound, err, t)
}

func TestRepository_UpdateUserPasswordErrorNoTable(t *testing.T) {
	mock.DropTables(db)

	err := repository.UpdateUserPassword(mock.GetFirstUser())
	utils.AssertNotNil(err, t)
}

func TestRepository_GetUserEmailSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	email, err := repository.GetUserEmail(1)
	utils.AssertNil(err, t)
	utils.AssertEqual(mock.GetFirstUser().Email, email, t)
}

func TestRepository_GetUserEmailUserNotFoundError(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	_, err := repository.GetUserEmail(mock.GetNotExistsUserId())
	utils.AssertErrorsEqual(internal_errors.UnableToFindUserById, err, t)
}

func TestRepository_GetUserEmailErrorNoTable(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetUserEmail(1)
	utils.AssertNotNil(err, t)
}
