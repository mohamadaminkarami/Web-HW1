package redis

import (
	"context"
	"github.com/mohamadaminkarami/Web-HW1/utils"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = initRedis()

func initRedis() *redis.Client {
	redisDB, err := strconv.Atoi(utils.GetEnv("redis_db", "0"))

	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("redis_addr", "localhost:6379"),
		Password: utils.GetEnv("redis_password", ""), // no password set
		DB:       redisDB,                            // use default DB
	})
}

func GetValueRedis(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func SetValueRedis(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
