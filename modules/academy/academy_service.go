package academy

import (
	"fmt"
	"net/http"
	netUrl "net/url"
	"strings"
	"sync"

	"goyo/libs/naver"
	"goyo/models"
	"goyo/models/academy"
	"goyo/modules/common"

	"github.com/PuerkitoBio/goquery"
)

type Serivce interface {
	// Service
	CrawlNaverPlaces()
	CrawlMobileNaverPlace()
	// Response
	NewGetListResponse(result []NaverBasicInfoDTO, total int64, query *GetListQuery) GetListResponse
}

type service struct{}

// ------------------- Service -------------------

func (service) CrawlNaverPlaces() {
	var adminiStrationDivisions []models.AdminiStrationDivision
	if err := common.GetCommonRepo().GetAdminiStrationDivision(&adminiStrationDivisions); err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {

		pageNo := 1
		if i == 1 {
			pageNo = 51
		}

		for _, v := range adminiStrationDivisions {
			// xx시 xx 군 xx구 요가
			searchKeyword := fmt.Sprintf("%s 요가", v.SiGunGu)
			query := &naver.PcPlacesQuery{
				Query:      searchKeyword,
				Start:      pageNo,
				Display:    50,
				Adult:      false,
				Spq:        false,
				QueryRank:  "",
				DeviceType: "pcmap",
			}

			var result naver.PcPlacesResult
			if err := naver.GetLib().GetPcPlaces(query, &result); err != nil {
				panic(err)
			}

			// 네이버 검색 결과 값을 DB에 저장
			for _, v := range result.Result.Items {

				// 요가원인지 판단하는 flag
				isRelatedYoga := v.Category == "요가원" || v.Category == "요가,명상"

				if !isRelatedYoga {
					continue
				}

				// 존재하는지 여부 체크
				isExist, err := GetRepo().FindNaverPlace(v.Id, nil)
				fmt.Println(isExist)
				if err != nil {
					panic(err)
				}

				if isExist == 1 {
					continue
				}

				newPlaceData := academy.NaverPlace{
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
				if err := GetRepo().CreateNaverPlace(&newPlaceData); err != nil {
					panic(err)
				}

			}
		}
	}
}

func (service) CrawlMobileNaverPlace() {
	var naverPlaces []CrawlMobileNaverPlaceDto
	if err := GetRepo().GetNaverId(&naverPlaces); err != nil {
		panic(err)
	}

	for _, v := range naverPlaces {

		// Access Naver Place
		url := fmt.Sprintf("https://m.place.naver.com/place/%s/home", v.NaverId)
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		// Read Html
		html, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}

		// Follow ThumbUrl Tag
		imageGroupDiv := html.Find("div.Y8J3x")
		firstImage := imageGroupDiv.Find("div.cb7hz")
		styleValue, _ := firstImage.Attr("style")

		// No Value
		if styleValue == "" {
			continue
		}

		// Slice https://~~~/"
		indexFromHttps := strings.Index(styleValue, "https")
		imageUrlWithQuotationMark := styleValue[indexFromHttps:]

		// Slice "
		indexFromQuotationMark := strings.Index(imageUrlWithQuotationMark, "\"")

		imgUrl := styleValue[:indexFromQuotationMark]
		decodedUrl, _ := netUrl.QueryUnescape(imgUrl)
		result := decodedUrl

		// DB update
		if result != "" {
			if err := GetRepo().UpdateNaverBasicInfo(v.Id, result); err != nil {
				continue
			}
		}
	}
}

// ------------------- Response -------------------

func (service) NewGetListResponse(result []NaverBasicInfoDTO, total int64, query *GetListQuery) GetListResponse {
	var response GetListResponse

	pageInfo := common.PaginationInfo{
		PageSize:  len(result),
		PageCount: total/int64(query.RowCount) + 1,
		Page:      int64(query.PageNo),
		RowCount:  total,
	}

	response.List = result
	response.Pagination = pageInfo

	return response
}

// ------------------- SingleTon -------------------

var (
	serviceInstance Serivce
	serviceOnce     sync.Once
)

func GetService() Serivce {
	if serviceInstance != nil {
		return serviceInstance
	}
	serviceOnce.Do(func() {
		serviceInstance = &service{}
	})
	return serviceInstance
}
