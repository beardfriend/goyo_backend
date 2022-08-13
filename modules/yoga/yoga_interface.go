package yoga

// --------- Response -----------

type GetSortsResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetQuery struct {
	Name string `form:"keyword" binding:"required"`
}

// ---------- DTO ---------------

type SortsDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (SortsDTO) TableName() string {
	return `yoga_sorts`
}
