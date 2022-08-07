package yoga

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

type YogaController struct{}

func (YogaController) GET(c *gin.Context) {
	query := new(GetQuery)

	if err := c.ShouldBindQuery(query); err != nil {
		common.SendError(c, 400, "쿼리스트링을 확인해주세요")
		return
	}

	isContainConsonants := false
	queryStringLength := utf8.RuneCountInString(query.Name)

	// queryString이 한 자일 때 체크
	if queryStringLength == 1 {
		// 자음인 경우
		if matched, _ := regexp.MatchString("[ㄱ-ㅎ]", query.Name); matched {
			isContainConsonants = true
		}
	}

	if queryStringLength > 1 {
		// 자음이 마지막에 있는지 체크
		if matched, _ := regexp.MatchString("[가-힣]+[가-힣ㄱ-ㅎ]$", query.Name); !matched {
			common.SendError(c, 400, "쿼리스트링 중간에 자음을 포함할 수 없습니다.")
			return
		}
		// 마지막이 자음일 때
		if matched, _ := regexp.MatchString("[ㄱ-ㅎ]$", query.Name); matched {
			isContainConsonants = true
		}
	}

	var result []YogaSorts

	if isContainConsonants {
		lastWord := fmt.Sprintf("%c", []rune(query.Name)[queryStringLength-1])
		searchWordFirst := map[string]string{"ㄱ": "가", "ㄴ": "나", "ㄷ": "다", "ㄹ": "라", "ㅁ": "마", "ㅂ": "바", "ㅅ": "사", "ㅇ": "아", "ㅈ": "자", "ㅊ": "차", "ㅋ": "카", "ㅌ": "타", "ㅍ": "파", "ㅎ": "하", "ㄲ": "까", "ㄸ": "따", "ㅆ": "싸", "ㅃ": "빠"}
		searchWordLast := map[string]string{"ㄱ": "낗", "ㄴ": "닣", "ㄷ": "딯", "ㄹ": "맇", "ㅁ": "밓", "ㅂ": "빟", "ㅅ": "싷", "ㅇ": "잏", "ㅈ": "짛", "ㅊ": "칳", "ㅋ": "킿", "ㅌ": "팋", "ㅍ": "핗", "ㅎ": "힣", "ㄲ": "낗", "ㄸ": "띻", "ㅆ": "앃", "ㅃ": "삫"}

		firstWord := strings.Trim(query.Name, lastWord) + searchWordFirst[lastWord]
		secondWord := strings.Trim(query.Name, lastWord) + searchWordLast[lastWord]

		if err := GetRepo().GetYogaSortByCosonants(firstWord, secondWord, &result); err != nil {
			panic(err)
		}
	} else {
		if err := GetRepo().GetYogaSort(query.Name, &result); err != nil {
			panic(err)
		}
	}

	response := make([]GetResponse, 0)

	for i, v := range result {
		response = append(response, GetResponse{Id: uint(i + 1), Name: v.Name})
	}

	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}
