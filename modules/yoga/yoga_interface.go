package yoga

import "goyo/models"

// ------------------- Request -------------------
type InsertSortsBody struct {
	Value []InsertYogaSortsBodyDetail `json:"value"`
}

type InsertYogaSortsBodyDetail struct {
	NaverPlaceID int    `json:"naverPlaceId" binding:"required"`
	Name         string `json:"name"`
}

// ------------------- Response -------------------
type GetSortsResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetQuery struct {
	Name string `form:"keyword" binding:"required"`
}

// ------------------- DTO -------------------

type SortsDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateSortsDTO struct {
	models.Model
	NaverPlaceId int    `json:"naver_place_id"`
	Name         string `json:"name"`
}

func (CreateSortsDTO) TableName() string {
	return `yoga_sorts`
}

func (SortsDTO) TableName() string {
	return `yoga_sorts`
}
