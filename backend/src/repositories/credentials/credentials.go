package credentials

import (
	"github.com/jmoiron/sqlx"
	"internal_errors"
	"models"
)

const (
	AddUserQuery             = `INSERT INTO users(email, password) VALUES(:email, :password)`
	UserIdByCredentialsQuery = `SELECT id FROM users WHERE email = :email AND password = :password`
	UserEmailByIdQuery       = `SELECT email FROM users WHERE id = $1`
	UpdateUserPasswordQuery  = `UPDATE users SET password = :password WHERE email = :email`
)

var (
	emailExistsError = `pq: duplicate key value violates unique constraint "users_email_key"`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) CreateUser(user models.UserCredentials) error {
	_, err := r.db.NamedExec(AddUserQuery, user)

	if err != nil && err.Error() == emailExistsError {
		err = internal_errors.UnableToRegisterUserEmailExists
	}

	return err
}

func (r Repository) GetUserIdByCredentials(user models.UserCredentials) (uint, error) {
	var id uint
	rows, err := r.db.NamedQuery(UserIdByCredentialsQuery, user)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&id)
	} else {
		return 0, internal_errors.UnableToLoginUserNotFound
	}

	return id, err
}

func (r Repository) UpdateUserPassword(user models.UserCredentials) error {
	res, err := r.db.NamedExec(UpdateUserPasswordQuery, user)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affectedRows == 0 {
		return internal_errors.UnableToChangePasswordUserNotFound
	}

	return nil
}

func (r Repository) GetUserEmail(userId uint) (string, error) {
	var email string
	rows, err := r.db.Query(UserEmailByIdQuery, userId)
	if err != nil {
		return "", err
	}

	if rows.Next() {
		err = rows.Scan(&email)
		return email, err
	} else {
		return "", internal_errors.UnableToFindUserById
	}
}
