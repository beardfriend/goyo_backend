package yoga

// --------- Response -----------

type GetResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetQuery struct {
	Name string `form:"keyword" binding:"required"`
}

// ---------- DTO ---------------

type YogaSorts struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
