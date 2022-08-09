package naver

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/machinebox/graphql"
)

type NaverPlaceQuery struct {
	Query      string `json:"query"`
	Start      int    `json:"start"`
	Display    int    `json:"display"`
	Adult      bool   `json:"adult"`
	Spq        bool   `json:"spq"`
	QueryRank  string `json:"queryRank"`
	DeviceType string `json:"deviceType"`
	X          string `json:"x"`
	Y          string `json:"y"`
}

type NaverPlaceResult struct {
	Result NaverPlaceItem `json:"result"`
}

type NaverPlaceItem struct {
	Items []NaverPlaceItemDetail `json:"items"`
}

type NaverPlaceItemDetail struct {
	Id            string  `json:"id"`       // 가게 아이디
	Name          string  `json:"name"`     // 가게 이름
	Category      string  `json:"category"` // 업종 카테고리
	Distance      string  `json:"distance"` // 현 위치로부터 거리
	RoadAddress   string  `json:"roadAddress"`
	Address       string  `json:"address"`
	FullAddress   string  `json:"fullAddress"`
	CommonAddress string  `json:"commonAddress"`
	BookingUrl    *string `json:"bookingUrl"`
	Phone         *string `json:"phone"`
	VirtualPhone  string  `json:"virtualPhone"`
	BusinessHours *string `json:"businessHours"`
	ImageUrl      *string `json:"imageUrl"`
	X             string  `json:"x"`
	Y             string  `json:"y"`
}

func (lib) MobileGet(naverId string) *resty.Response {
	url := fmt.Sprintf("https://m.place.naver.com/place/%s/home", naverId)
	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		Get(url)

	return resp
}

func (lib) Get(query *NaverPlaceQuery, result *NaverPlaceResult) error {
	client := graphql.NewClient("https://pcmap-api.place.naver.com/place/graphql")
	req := graphql.NewRequest(`
	query getPlacesList($input: PlacesInput) {
		result: places(input: $input) {
		  items {
			id
			name
			normalizedName
			category
			distance
			roadAddress
			address
			fullAddress
			commonAddress
			bookingUrl
			phone
			virtualPhone
			businessHours
			imageUrl
			x
			y
		  }
		}
	  }
	  
	`)

	req.Var("input", &query)
	ctx := context.Background()
	err := client.Run(ctx, req, &result)
	return err
}
