package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"urlshortener/utils"
)

func InsertLongToShortUrlCache(longUrl, shortUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := redisClient.Set(ctx, fmt.Sprintf("long:%s", longUrl), shortUrl, 0).Err()
	if err != nil {
		return utils.SettingCacheError
	}
	return nil
}

func InsertShortToLongUrlCache(longUrl, shortUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := redisClient.Set(ctx, fmt.Sprintf("short:%s", shortUrl), longUrl, 0).Err()
	if err != nil {
		return utils.SettingCacheError
	}
	return nil
}
