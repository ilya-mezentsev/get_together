package credentials

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"internal_errors"
	"models"
)

const (
	AddUserQuery = `
	INSERT INTO users(created_at)
	SELECT CURRENT_TIMESTAMP
	WHERE NOT EXISTS(
		SELECT 1 FROM users_credentials WHERE email = $1
	) RETURNING id`
	AddUserCredentialsQuery = `
	INSERT INTO users_credentials(user_id, email, password)
	VALUES(:user_id, :email, :password)`
	UserIdByCredentialsQuery = `
		SELECT user_id FROM users_credentials WHERE email = :email AND password = :password`
	UserEmailByIdQuery      = `SELECT email FROM users_credentials WHERE user_id = $1`
	UpdateUserPasswordQuery = `UPDATE users_credentials SET password = :password WHERE email = :email`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) CreateUser(user models.UserCredentials) error {
	var insertedUserId uint
	err := r.db.Get(&insertedUserId, AddUserQuery, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal_errors.UnableToRegisterUserEmailExists
		}

		return err
	}

	_, err = r.db.NamedExec(AddUserCredentialsQuery, map[string]interface{}{
		"user_id":  insertedUserId,
		"email":    user.Email,
		"password": user.Password,
	})
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
