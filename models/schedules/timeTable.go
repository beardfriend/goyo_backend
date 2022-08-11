package schedules

import (
	"goyo/libs/types"
)

type TimeTable struct {
	Id           uint             `gorm:"primaryKey"`
	NaverPlaceID uint             `gorm:"NOT NULL"`
	Sort         string           `gorm:"column:sort; VARCHAR; NOT NULL; size:100; comment: 요가 종류"`
	Year         string           `gorm:"column:year VARCHAR; NOT NULL; size:100; comment: 년"`
	Month        string           `gorm:"column:month VARCHAR; NOT NULL; size:100; comment: 월"`
	Day          string           `gorm:"column:day; VARCHAR; NOT NULL; size:20; comment: 요일"`
	StartTime    types.TimeString `gorm:"column:start_time; TIME; NOT NULL; comment: 요가 시작 시간"`
	EndTime      types.TimeString `gorm:"column:end_time; TIME; NOT NULL; comment: 요가 종료 시간"`
	YogaName     string           `gorm:"column:yoga_name; VARCHAR; size:256; comment: 요가 이름"`
	RoomName     string           `gorm:"column:room_name; VARCHAR; size:256; comment: 강의실 이름"`
	Level        string           `gorm:"column:level; VARCHAR; size:256; comment: 강의 난이도"`
	TeacherName  string           `gorm:"column:teacher_name; VARCHAR; size:256; comment: 선생님 이름"`
}

func (TimeTable) TableName() string {
	return "schedule_time_table"
}
