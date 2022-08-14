package admin

// ------------------- DTO -------------------

type NaverPlaceDTO struct {
	Id            uint          `json:"id"`
	IsRegist      bool          `json:"is_regist"`
	NaverId       string        `json:"naver_id"`
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
	YogaSorts     []YogaSortDTO `json:"yogaSorts" gorm:"foreignKey:NaverPlaceID"`
}

type YogaSortDTO struct {
	NaverPlaceID uint   `json:"id"`
	Name         string `json:"name"`
}

func (YogaSortDTO) TableName() string {
	return `yoga_sorts`
}
