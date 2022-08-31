package models

type Health struct {
	Primary
	Status string `gorm:"comment:상태"`
	TimeWithDeleted
}
