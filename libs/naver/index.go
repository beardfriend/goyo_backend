package naver

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

type Lib interface {
	Get(query *NaverPlaceQuery, result *NaverPlaceResult) error
	MobileGet(naverId string) *resty.Response
}

type lib struct{}

var (
	naverlibInstance Lib
	once             sync.Once
)

func GetLib() Lib {
	if naverlibInstance != nil {
		return naverlibInstance
	}
	once.Do(func() {
		naverlibInstance = &lib{}
	})
	return naverlibInstance
}
