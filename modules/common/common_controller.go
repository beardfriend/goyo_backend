package common

import (
	"goyo/models"

	"github.com/gin-gonic/gin"
)

type CommonController struct{}

//	지역구 리스트 조회
func (CommonController) GetAdministrations(c *gin.Context) {
	var result []models.AdminiStrationDivision
	GetCommonRepo().GetAdminiStrationDivision(&result)

	response := new(AdminiStrationsResponse)
	response.List = result

	SendResult(c, 200, "ok", response)
}
