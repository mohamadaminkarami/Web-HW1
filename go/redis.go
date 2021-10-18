package main

import (
	"github.com/go-redis/redis/v8"
)

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func getValueRedis(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func setValueRedis(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
