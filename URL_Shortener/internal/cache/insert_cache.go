package cache

import (
	"context"
	"fmt"
	"time"
	"urlshortener/utils"
)

func InsertLongToShortUrlCache(longUrl, shortUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := RedisClient.Set(ctx, fmt.Sprintf("long:%s", longUrl), shortUrl, 24*time.Hour).Err()
	if err != nil {
		return utils.SettingCacheError
	}
	return nil
}

func InsertShortToLongUrlCache(longUrl, shortUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := RedisClient.Set(ctx, fmt.Sprintf("short:%s", shortUrl), longUrl, 24*time.Hour).Err()
	if err != nil {
		return utils.SettingCacheError
	}
	return nil
}
