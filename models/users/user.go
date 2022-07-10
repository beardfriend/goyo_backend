package users

import (
	"goyo/libs/types"
	"goyo/models"
)

type User struct {
	models.Model
	NickName    string           `gorm:"size:20; not null"`
	PhoneNumber string           `gorm:"size:15"`
	LastLoginAt types.TimeString `gorm:"type:DATETIME"`
	IsAdmin     bool             `gorm:"not null; default:0"`
}
