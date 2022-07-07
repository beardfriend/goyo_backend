package users

import (
	"goyo/libs/types"
	"goyo/models"
)

type User struct {
	models.Model
	Email       string           `gorm:"size:100"`
	Password    string           `gorm:"size:100"`
	NickName    string           `gorm:"size:20"`
	PhoneNumber string           `gorm:"size:15"`
	LastLoginAt types.TimeString `gorm:"type:DATETIME"`
	IsAdmin     bool             `gorm:"not null; default:0"`
}
