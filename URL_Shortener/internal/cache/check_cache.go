package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"urlshortener/utils"
)

func CheckCacheForEncoding(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := redisClient.Get(ctx, fmt.Sprintf("long:%s", url)).Result()
	if err != nil {
		return "", utils.UnitNotFoundError
	}
	return val, nil
}

func CheckCacheForRedirect(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := redisClient.Get(ctx, fmt.Sprintf("short:%s", url)).Result()
	if err != nil {
		return "", utils.UnitNotFoundError
	}
	return val, err
}
