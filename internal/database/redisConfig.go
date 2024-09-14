package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/oik17/mpl-be/internal/utils"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client 

func RedisConnect() {
	redisDB, err := strconv.Atoi(utils.Config("REDIS_DB"))
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.Config("REDIS_HOST"), utils.Config("REDIS_PORT")),
		Password: utils.Config("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	ctx := context.Background()

	// Testing the Redis connection
	err = RedisClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	log.Println("foo", val)
	log.Println("Redis Connected")
}
