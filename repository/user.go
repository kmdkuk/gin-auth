package repository

import (
	"github.com/kmdkuk/gin-auth/db"
	"github.com/kmdkuk/gin-auth/model"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByID(userID string) (model.User, error) {
	d := db.GetDB()
	var user model.User

	if err := d.First(&user, "user_id=?", userID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func CreateUser(user model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	d := db.GetDB()
	d.NewRecord(user)
	d.Create(&user)
	if d.Error != nil {
		return d.Error
	}
	return nil
}
