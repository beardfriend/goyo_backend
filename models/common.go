package models

import "goyo/libs/types"

type Model struct {
	ID        uint             `gorm:"primarykey"`
	CreatedAt types.TimeString `gorm:"type:DATETIME"`
	UpdatedAt types.TimeString `gorm:"type:DATETIME"`
}
