package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	connectionString := os.Getenv("REDIS_CONNECTION_STRING")
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
