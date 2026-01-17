package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

const ADDR = "localhost:6379"
const PASSWORD = ""

type RedisAPI interface {
	SetKV(key string, value string, exp string) error
	GetKV(key string) (string, error)
}

var redisClient *redis.Client
var once sync.Once

// Init 初始化 Redis 客户端
func Init(addr string, password string) {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
		})
	})
}
