package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/mohamadaminkarami/Web-HW1/utils"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = initRedis()

func initRedis() *redis.Client {
	redisDB, err := strconv.Atoi(utils.GetEnv("REDIS_DB", "0"))
	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: utils.GetEnv("REDIS_PASSWORD", ""), // no password set
		DB:       redisDB,                            // use default DB = 0
	})
}

func PingRedis() {
	_, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println("connected to redis successfully")
	} else {
		log.Panic("Cannot connect to redis")
	}
}

func GetValueRedis(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func SetValueRedis(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
