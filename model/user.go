package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"unique"`
	Password string `json:"password"`
}
