package academy

import (
	"goyo/modules/common"
)

// ------------------- Body -------------------
type RegistRequestBody struct {
	Year        string `json:"year" binding:"required"`
	Month       string `json:"month" binding:"required"`
	Day         string `json:"day" binding:"required"`
	YogaName    string `json:"yogaName" binding:"required"`
	RoomName    string `json:"roomName"`
	StartTime   string `json:"startTime" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"`
	Level       string `json:"level"`
	TeacherName string `json:"teacherName"`
}

// ------------------- Query -------------------
type GetListQuery struct {
	YogaSort string `form:"yoga_sort"`
	SiGunGu  string `form:"si_gun_gu"`
	PageNo   int    `form:"page_no,default=1"`
	RowCount int    `form:"row_count,default=10"`
}

// -------------- Response ---------------------
type GetListResponse struct {
	List       []NaverPlaceDTO       `json:"list"`
	Pagination common.PaginationInfo `json:"pagination"`
}

// --------------- DTO ---------------------

type NaverPlaceDTO struct {
	Id            uint          `json:"id"`
	Name          string        `json:"name"`
	Category      string        `json:"category"`
	RoadAddress   string        `json:"roadAddress"`
	CommonAddress string        `json:"commonAddress"`
	BookingUrl    *string       `json:"bookingUrl"`
	PhoneNum      *string       `json:"phoneNum"`
	BusinessHours *string       `json:"businessHours"`
	ImageUrl      *string       `json:"imageUrl"`
	ThumbUrl      *string       `json:"thumbUrl"`
	X             string        `json:"x"`
	Y             string        `json:"y"`
	YogaSorts     []YogaSortDTO `json:"yogaSorts" gorm:"foreignKey:NaverBasicInfoID"`
}

type CrawlMobileNaverPlaceDto struct {
	Id       uint    `json:"id"`
	NaverId  string  `json:"naver_id"`
	ThumbUrl *string `json:"thumb_url,omitempty"`
}

type YogaSortDTO struct {
	NaverBasicInfoID uint   `json:"-"`
	Name             string `json:"name"`
}

func (YogaSortDTO) TableName() string {
	return `yoga_sort`
}
