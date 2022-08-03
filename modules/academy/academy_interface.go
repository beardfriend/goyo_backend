package academy

import (
	"goyo/modules/common"
	"goyo/modules/yoga"
)

// --------------- Request -------------------------
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

type GetListQuery struct {
	YogaSort string `form:"yoga_sort"`
	SiGunGu  string `form:"si_gun_gu"`
	PageNo   int    `form:"page_no,default=1"`
	RowCount int    `form:"row_count,default=10"`
}

// -------------- Response ---------------------
type GetListResponse struct {
	List       []GetListDetail       `json:"list"`
	Pagination common.PaginationInfo `json:"pagination"`
}

type GetListDetail struct {
	Id            uint            `json:"id"`
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	RoadAddress   string          `json:"roadAddress"`
	CommonAddress string          `json:"commonAddress"`
	BookingUrl    *string         `json:"bookingUrl"`
	PhoneNum      *string         `json:"phoneNum"`
	BusinessHours *string         `json:"businessHours"`
	ImageUrl      *string         `json:"imageUrl"`
	X             string          `json:"x"`
	Y             string          `json:"y"`
	YogaSort      []yoga.Response `json:"yogaSort"`
}

// --------------- DAO ---------------------

type NaverBasicInfoDAO struct {
	Id            uint
	Name          string
	Category      string
	RoadAddress   string
	CommonAddress string
	BookingUrl    *string
	PhoneNum      *string
	BusinessHours *string
	ImageUrl      *string
	X             string
	Y             string
}
