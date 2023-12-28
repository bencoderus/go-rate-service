package redis

import (
	"os"

	"github.com/gofiber/storage/redis/v2"
)

func GetRedisConnection() *redis.Storage {
	return redis.New(redis.Config{
		URL:   os.Getenv("REDIS_CONNECTION_STRING"),
		Reset: false,
	})

}
