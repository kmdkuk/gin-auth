package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kmdkuk/gin-auth/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByID(userID string) (model.User, error)
	CreateUser(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUserByID(userID string) (model.User, error) {
	var user model.User

	if err := u.db.First(&user, "user_id=?", userID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) CreateUser(user model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	u.db.NewRecord(user)
	u.db.Create(&user)
	if u.db.Error != nil {
		return u.db.Error
	}
	return nil
}
