package admin

import (
	"strconv"
	"strings"

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
	var total int64

	GetRepo().GetListThatHasnTagTotal(query, &total)

	result := make([]NaverPlaceDTO, 0)
	GetRepo().GetListThatHasntTag(query, &result)
	common.SendResult(c, 200, "ok", gin.H{"list": result, "total": total})
}

func (AdminController) GetAcademyDetail(c *gin.Context) {
	idParam := c.Param("naver_id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	var result GetDetailNaverPlaceDTO
	GetRepo().GetDetail(&id, &result)

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

func (AdminController) DeleteTag(c *gin.Context) {
	idParam := c.Param("naver_id")

	result := strings.Split(idParam, ",")
	var idList []uint64
	for _, v := range result {
		id, _ := strconv.ParseUint(v, 10, 64)
		idList = append(idList, id)
	}

	if err := GetRepo().DeleteSorts(&idList); err != nil {
		panic(err)
	}

	common.SendOk(c, 200, "ok")
}
