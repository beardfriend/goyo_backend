package yoga

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"goyo/libs"
	"goyo/modules/common"
	rd "goyo/server/redis"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

type YogaController struct{}

func (YogaController) GetSorts(c *gin.Context) {
	// Get Query And Validate
	query := new(GetQuery)
	if err := c.ShouldBindQuery(query); err != nil {
		common.SendError(c, 400, "쿼리스트링을 확인해주세요")
		return
	}

	// Declare
	nameLength := utf8.RuneCountInString(query.Name)
	name := query.Name

	// Get isContainConsoonants
	location := GetService().CheckContainConsonants(nameLength, name)

	// Pass
	if location == "middle" {
		response := make([]GetSortsResponse, 0)
		common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
		return
	}

	// Get Sorts
	var result []SortsDTO
	if location == "last" {
		from, to := GetService().MakeSearchKeywordForConsonants(nameLength, name)

		if err := GetRepo().GetSortsByCosonants(from, to, &result); err != nil {
			panic(err)
		}
	}
	if location == "no" {
		if err := GetRepo().GetSortsByName(name, &result); err != nil {
			panic(err)
		}
	}

	// Make Response
	response := GetService().NewSortsList(result)

	// Send
	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}

func (YogaController) GetSortsV2(c *gin.Context) {
	query := new(GetQuery)
	if err := c.ShouldBindQuery(query); err != nil {
		common.SendError(c, 400, "쿼리스트링을 확인해주세요")
		return
	}
	// 최대 카운트를 조회해서,
	count := rd.GetInstance().ZCount(c, query.Name, "1", "+inf").Val()
	if count > 0 {
		key, _ := rd.GetInstance().ZRevRange(c, query.Name, 0, count-1).Result()
	}

	// 토탈 카운트에서 상위 2개를 넣어준다. 나머지는 랜덤으로 밑에 카운트 낮은 것들을 돌린다.

	// key, _ := rd.GetInstance().ZRangeByLex(c, query.Name, &redis.ZRangeBy{Min: "-", Max: "+", Offset: 0, Count: 100}).Result()
	// key, _, _ := rd.GetInstance().ZScan(c, query.Name, uint64(0), "", int64(10)).Result()

	var value []SortsDTO
	for _, v := range key {
		value = append(value, SortsDTO{Name: v})
	}
	// Make Response
	response := GetService().NewSortsList(value)

	// Send
	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}

func (YogaController) CronYogaSorts(c *gin.Context) {
	var yogaSort []SortsDTO
	if err := GetRepo().GetYogaSortDistinct(&yogaSort); err != nil {
		panic(err)
	}

	fmt.Println(yogaSort)
	for _, a := range yogaSort {
		i := 0
		var before []string
		s := strings.TrimSpace(a.Name)
		for _, v := range s {

			temp := v - libs.FirstWord
			cho := temp / 588
			jong := temp % 28
			choWord := libs.HangulCHO[cho]
			beforeString := strings.Join(before, "")
			if jong != 0 {
				rd.GetInstance().ZAdd(c, beforeString+string(v-jong), redis.Z{Member: s})
			}

			if i == 0 {
				rd.GetInstance().ZAdd(c, choWord, redis.Z{Member: s})
				rd.GetInstance().ZAdd(c, string(v), redis.Z{Member: s})
				before = append(before, string(v))
				i++
				continue
			}

			if i == utf8.RuneCountInString(s)-1 {

				rd.GetInstance().ZAdd(c, beforeString+choWord, redis.Z{Member: s})
				rd.GetInstance().ZAdd(c, beforeString, redis.Z{Member: s})
				rd.GetInstance().ZAdd(c, s, redis.Z{Member: s})
			}

			if i > 0 && i < utf8.RuneCountInString(s)-1 {
				fmt.Println(beforeString)
				rd.GetInstance().ZAdd(c, beforeString, redis.Z{Member: s})
				rd.GetInstance().ZAdd(c, beforeString+choWord, redis.Z{Member: s})
			}
			before = append(before, string(v))
			i++
		}
	}
	common.SendOk(c, 200, "ok")
}
