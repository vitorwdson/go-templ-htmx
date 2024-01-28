package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type RouteHandler func(http.ResponseWriter, *http.Request) error

type server struct {
	DB      *sql.DB
	Redis   *redis.Client
	Logger  *log.Logger
	DevMode bool
}

func NewServer(db *sql.DB, r *redis.Client, logger *log.Logger, devMode bool) server {
	return server{
		DB:      db,
		Redis:   r,
		Logger:  logger,
		DevMode: devMode,
	}
}

func (s server) SetupRoutes() {
	if s.DevMode {
		fs := http.FileServer(http.Dir("./static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}
}
