package redisdb

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piqba/mtss-cli/pkg/logger"
	"github.com/piqba/mtss-go/pkg/errors"
)

var (
	ErrRedisDbCheckConn = errors.NewError("Redis: Fail to check connection")
)

func GetRedisDbClient() *redis.Client {

	clientInstance := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URI"),  // use default Addr
		Password:     os.Getenv("REDIS_PASS"), // no password set
		DB:           0,
		DialTimeout:  60 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	})

	_, err := clientInstance.Ping(context.TODO()).Result()
	if err != nil {
		logger.LogError(ErrRedisDbCheckConn.Error())
	}
	return clientInstance
}
