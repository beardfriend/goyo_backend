package yoga

import "goyo/models"

type YogaSort struct {
	models.Model
	NaverBasicInfoID uint   `gorm:"column:naver_basic_info_id; NOT NULL"`
	Name             string `gorm:"column:name; NOT NULL"`
	Level            int8   `gorm:"column:level;"`
}

func (YogaSort) TableName() string {
	return `yoga_sort`
}
