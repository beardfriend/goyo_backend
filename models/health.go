package models

type Health struct {
	Model
	Status string `gorm:"comment:상태"`
}
