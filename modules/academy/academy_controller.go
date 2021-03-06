package academy

import (
	"fmt"

	"goyo/libs/naver"
	"goyo/models"
	"goyo/models/academy"
	"goyo/modules/common"

	"github.com/labstack/echo/v4"
)

type AcademyController struct{}

func (AcademyController) CrawlNaver(c echo.Context) error {
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
			if err := naver.GetNaverLib().Get(&q, &result); err != nil {
				panic(err)
			}
			for _, v := range result.Result.Items {

				isRelatedYoga := v.Category == "요가원" || v.Category == "요가,명상"

				if !isRelatedYoga {
					continue
				}

				isExist := GetAcademyRepo().FindNaverBasicInfo(v.Id) == 1

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
				if err := GetAcademyRepo().InsertNaverBasicInfo(&insertValue); err != nil {
					panic(err)
				}

			}

		}
	}
	return c.JSON(200, "ok")
}
