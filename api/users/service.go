package users

import (
	"errors"
)

type UserService struct {
	UserRepo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (us *UserService) RegisterUser(user User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("username, password, and email are required")
	}

	existingUser, err := us.UserRepo.GetUserByUsername(user.Username)
	if err == nil && existingUser.Username != "" {
		return errors.New("username already exists")
	}

	err = us.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(username, password string) (User, error) {
	user, err := us.UserRepo.GetUserByUsername(username)
	if err != nil {
		return User{}, err
	}

	if user.Password != password {
		return User{}, errors.New("invalid username or password")
	}

	return user, nil
}
