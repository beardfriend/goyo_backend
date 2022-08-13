package yoga

import (
	"context"
	"fmt"
	"unicode/utf8"

	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

type YogaController struct{}

func (YogaController) InsertYoga(c *gin.Context) {
	ctx := context.Background()
	if err := GetService().SyncSearchKeyword(ctx, "아쉬탕가", "아쉬탕가"); err != nil {
		panic(err)
	}

	str, _ := GetService().GetSearchKeyword(ctx, "아")
	fmt.Println(str)

	common.SendOk(c, 200, "ok")
}

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
