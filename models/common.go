package models

import "goyo/libs/types"

type Model struct {
	ID        uint             `json:"id" gorm:"primarykey"`
	CreatedAt types.TimeString `json:"created_at" gorm:"type:DATETIME"`
	UpdatedAt types.TimeString `json:"updated_at" gorm:"type:DATETIME"`
}
