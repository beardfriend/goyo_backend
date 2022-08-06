package yoga

import (
	"fmt"
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
	}

	isContainConsonants := false

	consonants := map[string]string{"ㄱ": "ㄱ", "ㄴ": "ㄴ", "ㄷ": "ㄷ", "ㄹ": "ㄹ", "ㅁ": "ㅁ", "ㅂ": "ㅂ", "ㅅ": "ㅅ", "ㅇ": "ㅇ", "ㅈ": "ㅈ", "ㅊ": "ㅊ", "ㅋ": "ㅋ", "ㅌ": "ㅌ", "ㅍ": "ㅍ", "ㅎ": "ㅎ"}

	// queryString이 한 자일 때 체크
	if utf8.RuneCountInString(query.Name) == 1 {
		// 자음인 경우
		if len(consonants[string(query.Name[0])]) == 1 {
			isContainConsonants = true
		}
	}
	// queryString이 1자 이상일 때
	if utf8.RuneCountInString(query.Name) > 1 {
		// 자음이 마지막에 있는지 체크
		var lastW string
		for i := 0; i < utf8.RuneCountInString(query.Name)-1; i++ {
			lastW = fmt.Sprintf("%c", []rune(query.Name)[i])
			if len(consonants[lastW]) > 1 {
				common.SendError(c, 400, "쿼리스트링 중간에 자음을 포함할 수 없습니다.")
				return
			}
		}
		queryLen := utf8.RuneCountInString(query.Name)

		lastWw := fmt.Sprintf("%c", []rune(query.Name)[queryLen-1])

		// 자음이면
		if consonants[lastWw] == lastWw {
			isContainConsonants = true
		}
	}

	var result []YogaSorts

	if isContainConsonants {
		lastW := fmt.Sprintf("%c", []rune(query.Name)[utf8.RuneCountInString(query.Name)-1])
		searchWordFirst := map[string]string{"ㄱ": "가", "ㄴ": "나", "ㄷ": "다", "ㄹ": "라", "ㅁ": "마", "ㅂ": "바", "ㅅ": "사", "ㅇ": "아", "ㅈ": "자", "ㅊ": "차", "ㅋ": "카", "ㅌ": "타", "ㅍ": "파", "ㅎ": "하", "ㄲ": "까", "ㄸ": "따", "ㅆ": "싸", "ㅃ": "빠"}
		searchWordLast := map[string]string{"ㄱ": "낗", "ㄴ": "닣", "ㄷ": "딯", "ㄹ": "맇", "ㅁ": "밓", "ㅂ": "빟", "ㅅ": "싷", "ㅇ": "잏", "ㅈ": "짛", "ㅊ": "칳", "ㅋ": "킿", "ㅌ": "팋", "ㅍ": "핗", "ㅎ": "힣", "ㄲ": "낗", "ㄸ": "띻", "ㅆ": "앃", "ㅃ": "삫"}

		firstWord := strings.Trim(query.Name, lastW) + searchWordFirst[lastW]
		secondWord := strings.Trim(query.Name, lastW) + searchWordLast[lastW]

		if err := GetRepo().GetYogaSortByCosonants(firstWord, secondWord, &result); err != nil {
			panic(err)
		}
	} else {
		if err := GetRepo().GetYogaSort(query.Name, &result); err != nil {
			panic(err)
		}
	}

	var response []GetResponse

	for i, v := range result {
		response = append(response, GetResponse{Id: uint(i + 1), Name: v.Name})
	}

	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}
