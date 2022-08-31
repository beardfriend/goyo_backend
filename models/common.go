package models

import (
	"goyo/libs/types"

	"gorm.io/gorm"
)

type Primary struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type DefaultTime struct {
	CreatedAt types.TimeString `json:"created_at" gorm:"type:DATETIME; NOT NULL"`
	UpdatedAt types.TimeString `json:"updated_at" gorm:"type:DATETIME"`
}

type TimeWithDeleted struct {
	CreatedAt types.TimeString `json:"created_at" gorm:"type:DATETIME; NOT NULL"`
	UpdatedAt types.TimeString `json:"updated_at" gorm:"type:DATETIME"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}
