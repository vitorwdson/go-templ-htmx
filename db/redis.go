package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func MustConnectRedis() *redis.Client {
	connectionString := os.Getenv("REDIS_CONNECTION_STRING")
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	err = client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return client
}
