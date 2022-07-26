package academy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

type AcademyController struct{}

func (AcademyController) CrawlNaver(c echo.Context) error {
	search := "서울시 항동 요가학원"
	encodedUrl := url.QueryEscape(search)
	url := fmt.Sprintf("https://pcmap.place.naver.com/place/list?query=%s", encodedUrl)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	html, _ := goquery.NewDocumentFromReader(resp.Body)
	index := strings.Index(html.Text(), "{\"$ROOT_QUERY.places")
	index2 := strings.Index(html.Text(), "window.__PLACE_STATE__")
	fmt.Println(index, index2)
	str := strings.Trim(html.Text()[index:index2], " ")
	dd := str[:len(str)-2]
	var ss map[string]interface{}

	json.Unmarshal([]byte(dd), &ss)
	// var result []models.AdminiStrationDivision
	// mariadb.GetInstance().Raw(`SELECT si_gun_gu FROM administration_division WHERE si_do='서울'`).Find(&result)
	firstQuery := fmt.Sprintf("$ROOT_QUERY.places({\"input\":{\"adult\":false,\"deviceType\":\"pcmap\",\"display\":50,\"query\":\"%s\",\"queryRank\":\"\",\"spq\":false,\"start\":1}})", search)
	ff := ss[firstQuery].(map[string]interface{})
	ff2 := ff["items"]
	return c.JSON(200, ff2)
}
