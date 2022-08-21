package redis

import (
	"fmt"
	"sync"

	"goyo/libs"

	"github.com/go-redis/redis/v9"
)

var (
	instance  *redis.Client
	onceRedis sync.Once
)

func GetInstance() *redis.Client {
	if instance != nil {
		return instance
	}
	onceRedis.Do(func() {
		addr := fmt.Sprintf("%s:%s", libs.ENV.Redis.Host, libs.ENV.Redis.Port)

		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		})
		instance = rdb
	})
	return instance
}
