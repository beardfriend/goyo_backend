package academy

import (
	"fmt"
	"net/http"

	"goyo/libs/naver"
	"goyo/models"
	"goyo/models/academy"
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

type AcademyController struct{}

func (AcademyController) CrawlNaver(c *gin.Context) {
	var gu []models.AdminiStrationDivision
	if err := common.GetCommonRepo().Getgu(&gu); err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		startNum := 1
		if i == 1 {
			startNum = 51
		}
		for _, v := range gu {
			queryString := fmt.Sprintf("%s 요가", v.SiGunGu)
			q := naver.NaverPlaceQuery{Query: queryString, Start: startNum, Display: 50, Adult: false, Spq: false, QueryRank: "", DeviceType: "pcmap"}
			var result naver.NaverPlaceResult
			if err := naver.GetLib().Get(&q, &result); err != nil {
				panic(err)
			}
			for _, v := range result.Result.Items {

				isRelatedYoga := v.Category == "요가원" || v.Category == "요가,명상"

				if !isRelatedYoga {
					continue
				}

				isExist := GetRepo().FindNaverBasicInfo(v.Id) == 1

				if isExist {
					continue
				}

				insertValue := academy.NaverBasicInfo{
					NaverId:       v.Id,
					Name:          v.Name,
					Category:      v.Category,
					RoadAddress:   v.RoadAddress,
					CommonAddress: v.CommonAddress,
					BookingUrl:    v.BookingUrl,
					PhoneNum:      v.Phone,
					BusinessHours: v.BusinessHours,
					ImageUrl:      v.ImageUrl,
					X:             v.X,
					Y:             v.Y,
				}
				if err := GetRepo().InsertNaverBasicInfo(&insertValue); err != nil {
					panic(err)
				}

			}

		}
	}
	common.SendOk(c, http.StatusCreated, "ok")
}

func (AcademyController) GetList(c *gin.Context) {
	query := new(GetListQuery)

	if err := c.ShouldBindQuery(&query); err != nil {
		errorMessage := common.BindJsonError(err, "GetListQuery")
		common.SendError(c, 400, errorMessage+" 를 입력해주세요")
	}

	var academyList []NaverBasicInfoDTO
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
