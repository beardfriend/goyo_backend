package yoga

import (
	"context"
	"sync"

	"goyo/server/redis"
)

type Serivce interface {
	SyncSearchKeyword(c context.Context, key string, keyword string) error
	GetSearchKeyword(c context.Context, key string) ([]string, error)
}

type service struct{}

func (service) SyncSearchKeyword(c context.Context, key string, keyword string) error {
	return redis.GetInstance().Set(c, key, keyword, 0).Err()
}

func (service) GetSearchKeyword(c context.Context, key string) ([]string, error) {
	return redis.GetInstance().Keys(c, "*"+key+"*").Result()
}

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
