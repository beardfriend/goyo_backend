package yoga

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"goyo/server/redis"
)

type Serivce interface {
	SyncSearchKeyword(c context.Context, key string, keyword string) error
	GetSearchKeyword(c context.Context, key string) ([]string, error)
	CheckContainConsonants(nameLength int, name string) string
	MakeSearchKeywordForConsonants(nameLength int, name string) (string, string)
	NewSortsList(value []SortsDTO) []GetSortsResponse
}

type service struct{}

// ------------------- Test -------------------

func (service) SyncSearchKeyword(c context.Context, key string, keyword string) error {
	return redis.GetInstance().Set(c, key, keyword, 0).Err()
}

func (service) GetSearchKeyword(c context.Context, key string) ([]string, error) {
	return redis.GetInstance().Keys(c, "*"+key+"*").Result()
}

// ------------------- Service -------------------

func (service) CheckContainConsonants(nameLength int, name string) string {
	location := "no"

	// queryString이 한 자일 때 체크
	if nameLength == 1 {
		// 자음인 경우
		if matched, _ := regexp.MatchString("[ㄱ-ㅎ]", name); matched {
			location = "last"
		}
	}

	if nameLength > 1 {
		// 쿼리스트링 중간에 자음이 있는지 체크
		if matched, _ := regexp.MatchString("[가-힣]+[가-힣ㄱ-ㅎ]$", name); !matched {
			location = "middle"
		}
		// 마지막이 자음일 때
		if matched, _ := regexp.MatchString("[ㄱ-ㅎ]$", name); matched {
			location = "last"
		}
	}
	return location
}

func (service) MakeSearchKeywordForConsonants(nameLength int, name string) (string, string) {
	lastWord := fmt.Sprintf("%c", []rune(name)[nameLength-1])

	fromArr := map[string]string{"ㄱ": "가", "ㄴ": "나", "ㄷ": "다", "ㄹ": "라", "ㅁ": "마", "ㅂ": "바", "ㅅ": "사", "ㅇ": "아", "ㅈ": "자", "ㅊ": "차", "ㅋ": "카", "ㅌ": "타", "ㅍ": "파", "ㅎ": "하", "ㄲ": "까", "ㄸ": "따", "ㅆ": "싸", "ㅃ": "빠"}
	toArr := map[string]string{"ㄱ": "낗", "ㄴ": "닣", "ㄷ": "딯", "ㄹ": "맇", "ㅁ": "밓", "ㅂ": "빟", "ㅅ": "싷", "ㅇ": "잏", "ㅈ": "짛", "ㅊ": "칳", "ㅋ": "킿", "ㅌ": "팋", "ㅍ": "핗", "ㅎ": "힣", "ㄲ": "낗", "ㄸ": "띻", "ㅆ": "앃", "ㅃ": "삫"}

	from := strings.Trim(name, lastWord) + fromArr[lastWord]
	to := strings.Trim(name, lastWord) + toArr[lastWord]

	return from, to
}

// ------------------- Repsonse -------------------

func (service) NewSortsList(value []SortsDTO) []GetSortsResponse {
	response := make([]GetSortsResponse, 0)
	for i, v := range value {
		response = append(response, GetSortsResponse{Id: uint(i + 1), Name: v.Name})
	}
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
