package database

import "github.com/go-redis/redis/v8"

func GetCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return rdb
}
