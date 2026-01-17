package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

func (api *RedisAPI) SetKV(key string, value string, exp string) error {
	return redisClient.Set(ctx, key, value, 0).Err()
}

func (api *RedisAPI) GetKV(key string) (string, error) {
	if redisClient == nil {
		
		return "", nil
	}
	return redisClient.Get(ctx, key).Result()
}