package user

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
