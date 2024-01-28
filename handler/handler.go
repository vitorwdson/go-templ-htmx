package handler

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type handler struct {
	DB      *sql.DB
	Redis   *redis.Client
	DevMode bool
}

func New(db *sql.DB, r *redis.Client, devMode bool) handler {
	return handler{
		DB:      db,
		Redis:   r,
		DevMode: devMode,
	}
}

func (h handler) SetupRoutes() {
}
