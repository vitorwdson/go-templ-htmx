package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type RouteHandler func(http.ResponseWriter, *http.Request) error

type handler struct {
	DB      *sql.DB
	Redis   *redis.Client
	Logger  *log.Logger
	DevMode bool
}

func New(db *sql.DB, r *redis.Client, logger *log.Logger, devMode bool) handler {
	return handler{
		DB:      db,
		Redis:   r,
		Logger:  logger,
		DevMode: devMode,
	}
}

func (h handler) SetupRoutes() {
}
