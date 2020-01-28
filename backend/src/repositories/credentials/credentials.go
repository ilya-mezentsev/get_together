package credentials

import (
  "github.com/jmoiron/sqlx"
  "internal_errors"
  "models"
)

const (
  AddUserQuery = `INSERT INTO users(email, password) VALUES(:email, :password)`
  UserIdByCredentialsQuery = `SELECT id FROM users WHERE email = :email AND password = :password`
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
