package store

import (
	"mini-wallet/libs"

	"github.com/go-redis/redis/v8"
)

func InitRedis(config libs.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "",
		DB:       0,
	})
	return rdb
}
