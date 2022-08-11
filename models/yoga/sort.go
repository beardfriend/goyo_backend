package yoga

import "goyo/models"

type YogaSort struct {
	models.Model
	NaverPlaceID uint   `gorm:"column:naver_place_id; NOT NULL"`
	Name         string `gorm:"column:name; NOT NULL"`
	Level        int8   `gorm:"column:level;"`
}

func (YogaSort) TableName() string {
	return `yoga_sort`
}
