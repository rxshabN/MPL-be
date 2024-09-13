package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/oik17/mpl-be/internal/utils"
	"github.com/redis/go-redis/v9"
)

func RedisConnect() {

	redisDB, err := strconv.Atoi(utils.Config("REDIS_DB"))
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.Config("REDIS_HOST"), utils.Config("REDIS_PORT")),
		Password: utils.Config("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	ctx := context.Background()

	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	log.Println("foo", val)
	log.Println("Redis Connected")
}
