package yoga

import "goyo/models"

type YogaScore struct {
	models.Primary
	Name  string `gorm:"column:name; NOT NULL; size:100;"`
	Score uint   `gorm:"column:score; NOT NULL;"`
	models.DefaultTime
}

func (YogaScore) TableName() string {
	return `yoga_score`
}
