package admin

import (
	"goyo/models"
	"goyo/modules/academy"
	"goyo/modules/common"
	"goyo/modules/yoga"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (AdminController) GetAcademyListThatHanstTag(c *gin.Context) {
	query := new(academy.AcademyListRequest)

	if err := c.ShouldBindQuery(query); err != nil {
		common.SendError(c, 400, "QueryString을 확인해주세요")
	}
	result := make([]academy.NaverPlaceDTO, 0)
	academy.GetRepo().GetListThatHasntTag(query, &result)
	common.SendResult(c, 200, "ok", result)
}

func (AdminController) InsertYogaSorts(c *gin.Context) {
	body := new(yoga.InsertSortsBody)

	if err := c.ShouldBindJSON(body); err != nil {
		common.SendError(c, 400, "Body를 확인해주세요")
	}

	DTO := yoga.GetService().NewSortsDTO(*body)

	if err := yoga.GetRepo().CreateSorts(&DTO); err != nil {
		panic(err)
	}

	common.SendOk(c, 201, "success")
}

func (AdminController) GetAdministrations(c *gin.Context) {
	var result []models.AdminiStrationDivision
	common.GetCommonRepo().GetAdminiStrationDivision(&result)

	response := new(common.AdminiStrationsResponse)
	response.List = result

	common.SendResult(c, 200, "ok", response)
}
