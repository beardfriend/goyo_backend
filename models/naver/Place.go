package naver

import (
	"goyo/models"
	"goyo/models/yoga"
)

// 서비스가 확장될 경우 회원가입을 받으면서 정보를 직접 받기로 ( 임시 테이블 )
type NaverPlace struct {
	models.Primary
	YogaSorts     []yoga.YogaSorts `gorm:"constraint:OnUpdate:CASCADE"`
	NaverId       string           `gorm:"index; column:naver_id; VARCHAR; NOT NULL; size:20; comment:네이버 고유 아이디"`
	Name          string           `gorm:"column:name; index; VARCHAR; NOT NULL; size:100; comment: 학원 이름"`
	Category      string           `gorm:"column:category; VARCHAR; NOT NULL; size:10; comment: 카테고리"`
	RoadAddress   string           `gorm:"column:road_address; VARCHAR; NOT NULL; size:100; comment: 도로명 주소"`
	CommonAddress string           `gorm:"column:common_address; VARCHAR; size:100; comment: 일반 주소"`
	BookingUrl    *string          `gorm:"column:booking_url; VARCHAR; size:256; comment: 예약 주소"`
	PhoneNum      *string          `gorm:"column:phone_num; VARCHAR; size:100; comment: 전화번호"`
	BusinessHours *string          `gorm:"column:business_hours; TEXT; comment: 영업시간"`
	ImageUrl      *string          `gorm:"column:image_url; VARCHAR; size:256; comment: 대표 이미지 주소"`
	ThumbUrl      *string          `gorm:"column:thumb_url; VARCHAR; size:1000; comment: 썸네일 이미지 주소"`
	X             string           `gorm:"column:x; VARCHAR; size:100; comment: x좌표"`
	Y             string           `gorm:"column:y; VARCHAR; size:100; comment: y좌표"`
	models.TimeWithDeleted
}

func (NaverPlace) TableName() string {
	return "naver_place"
}
