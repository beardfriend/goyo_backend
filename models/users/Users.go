package users

import (
	"goyo/libs/types"
	"goyo/models"
)

type Users struct {
	models.Model
	NickName    string           `gorm:"column:nickname; size:20; NOT NULL"`
	PhoneNumber string           `gorm:"column:phone_number; size:15"`
	LastLoginAt types.TimeString `gorm:"column:last_login_at; type:DATETIME"`
	RemovedAt   string           `gorm:"column:removed_at"`
	Tries       int8             `gorm:"column:tries; NOT NULL; default:0"`
	IsLock      bool             `gorm:"column:is_lock; NOT NULL; default:0"`
	IsAdmin     bool             `gorm:"column:is_admin; NOT NULL; default:0"`
	Type        int8             `gorm:"column:type; comment: 0 : academy 1: normal"`
}
