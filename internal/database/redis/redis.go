package redis

import (
	"fmt"
	"os"

	"github.com/gofiber/storage/redis/v2"
)

func GetRedisConnection() (redisClient *redis.Storage, err error) {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println(r)
			err = fmt.Errorf(fmt.Sprint(r))
		}
	}()

	redisClient = redis.New(redis.Config{
		URL:   os.Getenv("REDIS_CONNECTION_STRING"),
		Reset: false,
	})

	return redisClient, nil

}
