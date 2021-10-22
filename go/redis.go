package main

import (
	"strconv"

	"github.com/go-redis/redis/v8"
)

func initRedis() *redis.Client {
	redisDB, err := strconv.Atoi(getenv("redis_db", "0"))

	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     getenv("redis_addr", "localhost:6379"),
		Password: getenv("redis_password", ""), // no password set
		DB:       redisDB,                      // use default DB
	})
}

func getValueRedis(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func setValueRedis(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
