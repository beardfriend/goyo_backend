package naver

import "sync"

type Lib interface {
	Get(query *NaverPlaceQuery, result *NaverPlaceResult) error
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
