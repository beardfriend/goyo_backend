package yoga

import "goyo/models"

type YogaScore struct {
	models.Model
	Name  string `gorm:"column:name; NOT NULL; size:100;"`
	Score uint   `gorm:"column:score; NOT NULL;"`
}

func (YogaScore) TableName() string {
	return `yoga_score`
}
