package yoga

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"goyo/libs"
	"goyo/models/yoga"
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

	count, _ := rd.GetInstance().ZCount(c, query.Name, "0", "1").Result()

	var result []string
	var rankedResult []string
	var randomResult []string
	var uniqueResult []string
	if count == 1 {
		uniqueResult, _, _ = rd.GetInstance().ZScan(c, query.Name, 0, "", 1).Result()
	}

	if count == 2 {
		rankedIndex := 1
		uniqueResult, _, _ = rd.GetInstance().ZScan(c, query.Name, 0, query.Name, 1).Result()
		if len(uniqueResult) > 1 {
			rankedIndex = 0
		}
		rankedResult, _ = rd.GetInstance().ZRevRange(c, query.Name, 0, int64(rankedIndex)).Result()
	}

	if count > 2 {
		if count < 8 {
			fmt.Println(count)
			uniqueResult, _, _ = rd.GetInstance().ZScan(c, query.Name, 0, "", 0).Result()
			rankedResult, _ = rd.GetInstance().ZRevRange(c, query.Name, 0, count-2).Result()
		} else {
			rankedResult, _ = rd.GetInstance().ZRevRange(c, query.Name, 0, 1).Result()
			randomResult, _ = rd.GetInstance().ZRandMember(c, query.Name, 6).Result()
		}
	}
	if len(uniqueResult) > 1 {
		result = append(result, uniqueResult[0])
	}
	if len(rankedResult) > 0 {
		result = append(result, rankedResult...)
	}
	if len(randomResult) > 0 {
		result = append(result, randomResult...)
	}

	var value []SortsDTO
	for _, v := range result {
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

	for _, a := range yogaSort {
		uniqueResult, _, _ := rd.GetInstance().ZScan(c, a.Name, 0, "", 0).Result()

		if len(uniqueResult) > 1 {
			continue
		}

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

func (YogaController) UpdateScore(c *gin.Context) {
	body := new(UpdateScoreBody)

	if err := c.ShouldBindJSON(body); err != nil {
		common.SendError(c, 400, "잘못된 요청입니다.")
		return
	}

	arr, _, _ := rd.GetInstance().ZScan(c, body.Member, 0, body.Member, 0).Result()
	if len(arr) < 1 {
		common.SendError(c, 400, "없음")
		return
	}

	rd.GetInstance().ZIncrBy(c, body.Keyword, 1, body.Member)
	common.SendOk(c, 201, "ok")
}

func (YogaController) GetRanking(c *gin.Context) {
	var result []yoga.YogaScore
	if err := GetRepo().GetRanking(&result); err != nil {
		panic(err)
	}

	common.SendResult(c, 200, "ok", result)
}

func (YogaController) Ranking(c *gin.Context) {
	var yogaSort []SortsDTO
	if err := GetRepo().GetYogaSortDistinct(&yogaSort); err != nil {
		panic(err)
	}

	keys, _ := rd.GetInstance().Keys(c, "*").Result()
	result, _ := rd.GetInstance().ZUnionWithScores(c, redis.ZStore{Keys: keys}).Result()

	var searchValue []string
	scoreValue := make(map[string]*yoga.YogaScore, 0)

	for _, v := range result {
		if v.Score == 0 {
			continue
		}
		searchValue = append(searchValue, v.Member.(string))
		scoreValue[v.Member.(string)] = &yoga.YogaScore{Name: v.Member.(string), Score: uint(v.Score)}
	}

	var alreadyExist []yoga.YogaScore
	if err := GetRepo().GetScores(searchValue, &alreadyExist); err != nil {
		panic(err)
	}

	for _, v := range alreadyExist {
		fmt.Println(v.ID, scoreValue[v.Name].Score)
		GetRepo().UpdateCounts(v.ID, scoreValue[v.Name].Score)
		scoreValue[v.Name] = nil
	}

	var insertValue []yoga.YogaScore
	for _, v := range searchValue {
		if scoreValue[v] == nil {
			continue
		}
		insertValue = append(insertValue, yoga.YogaScore{Name: v, Score: scoreValue[v].Score})
	}

	if len(insertValue) > 0 {
		if err := GetRepo().CreateCounts(&insertValue); err != nil {
			panic(err)
		}
	}

	common.SendOk(c, 201, "ok")
}
