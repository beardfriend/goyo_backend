package academy

import (
	"net/http"

	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

type AcademyController struct{}

// 네이버 플레이스 정보 수집.
func (AcademyController) CrawlNaver(c *gin.Context) {
	go GetService().CrawlNaverPlaces()
	common.SendOk(c, http.StatusCreated, "ok")
}

// 썸네일 이미지 업데이트
func (AcademyController) UpdateThumbUrl(c *gin.Context) {
	go GetService().CrawlMobileNaverPlace()
	common.SendOk(c, http.StatusOK, "ok")
}

// 네이버 플레이스 학원 리스트 조회
func (AcademyController) GetList(c *gin.Context) {
	query := new(GetListQuery)
	if err := c.ShouldBindQuery(&query); err != nil {
		errorMessage := common.BindJsonError(err, "GetListQuery")
		common.SendError(c, 400, errorMessage+" 를 입력해주세요")
		return
	}

	var academyList []NaverPlaceDTO
	if err := GetRepo().GetAcademyListByYoga(query, &academyList); err != nil {
		panic(err)
	}

	var total int64
	if err := GetRepo().GetAcademyTotalByYoga(query, &total); err != nil {
		panic(err)
	}

	response := GetService().NewGetListResponse(academyList, total, query)

	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}
