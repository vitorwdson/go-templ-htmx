package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/vitorwdson/go-templ-htmx/data/repos"
)

type RouteHandler func(http.ResponseWriter, *http.Request) error

type server struct {
	DB       *sql.DB
	Redis    *redis.Client
	Logger   *log.Logger
	UserRepo *repos.UserRepo
	DevMode  bool
}

func NewServer(db *sql.DB, r *redis.Client, logger *log.Logger, devMode bool) server {
	userRepo := repos.NewUserRepo(db)

	return server{
		DB:       db,
		Redis:    r,
		Logger:   logger,
		DevMode:  devMode,
		UserRepo: &userRepo,
	}
}

func (s server) SetupRoutes() {
	http.HandleFunc("/register", s.handleErrors(s.handleRegister))

	if s.DevMode {
		fs := http.FileServer(http.Dir("./static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}
}
