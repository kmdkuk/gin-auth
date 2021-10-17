package repository

import (
	"fmt"

	"github.com/kmdkuk/gin-auth/model"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]model.User{}

func GetUserByID(userID string) (model.User, error) {
	// TODO: connect DB
	user, ok := users[userID]
	if !ok {
		return model.User{}, fmt.Errorf("not found user id: %s", userID)
	}
	return user, nil
}

func CreateUser(user model.User) error {
	if _, ok := users[user.ID]; ok {
		return fmt.Errorf("existing id %s", user.ID)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	// TODO: connect DB
	users[user.ID] = user
	return nil
}
