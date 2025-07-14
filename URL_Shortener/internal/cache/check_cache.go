package cache

import (
	"context"
	"fmt"
	"time"
	"urlshortener/utils"
)

func CheckCacheForEncoding(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	val, err := RedisClient.Get(ctx, fmt.Sprintf("long:%s", url)).Result()
	if err != nil {
		return "", utils.UnitNotFoundError
	}
	return val, nil
}

func CheckCacheForRedirect(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	val, err := RedisClient.Get(ctx, fmt.Sprintf("short:%s", url)).Result()
	if err != nil {
		return "", utils.UnitNotFoundError
	}
	return val, err
}
