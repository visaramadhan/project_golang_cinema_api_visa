package users

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user User) error {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
	_, err := ur.DB.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return errors.New("failed to create user: " + err.Error())
	}
	return nil
}

func (ur *UserRepository) GetUserByUsername(username string) (User, error) {
	query := "SELECT id, username, password, email FROM users WHERE username = $1"
	row := ur.DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	} else if err != nil {
		return User{}, errors.New("error retrieving user: " + err.Error())
	}
	return user, nil
}
