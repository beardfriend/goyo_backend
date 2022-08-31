package yoga

import "goyo/models"

type YogaSorts struct {
	models.Primary
	NaverPlaceID uint   `gorm:"column:naver_place_id; NOT NULL;"`
	ScheduleID   uint   `gorm:"column:schedule_id;"`
	Name         string `gorm:"column:name; size:100; NOT NULL"`
	Description  string `gorm:"column:description; VARCHAR; size:256"`
	Level        int8   `gorm:"column:level"`
	models.DefaultTime
}

func (YogaSorts) TableName() string {
	return `yoga_sorts`
}
