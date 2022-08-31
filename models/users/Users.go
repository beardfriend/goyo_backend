package users

import (
	"goyo/libs/types"
	"goyo/models"
)

type Users struct {
	models.Primary
	Email       string           `gorm:"column:email; size:100; NOT NULL"`
	NickName    string           `gorm:"column:nickname; size:20; NOT NULL"`
	PhoneNumber string           `gorm:"column:phone_number; size:15"`
	Type        int8             `gorm:"column:type; comment: 0 :academy 1:teacher 2:normal"`
	IsAdmin     bool             `gorm:"column:is_admin; NOT NULL; default:0"`
	IsLock      bool             `gorm:"column:is_lock; NOT NULL; default:0"`
	Tries       int8             `gorm:"column:tries; NOT NULL; default:0"`
	LastLoginAt types.TimeString `gorm:"column:last_login_at; type:DATETIME"`
	UserAcademy UserAcademy      `gorm:"foreignKey:ID"`
	models.TimeWithDeleted
}
