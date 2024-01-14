package user

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	DB    *sql.DB
	Redis *redis.Client
}
