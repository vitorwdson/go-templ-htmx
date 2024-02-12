package handler

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/vitorwdson/go-templ-htmx/db"
)

type RouteHandler func(http.ResponseWriter, *http.Request) error

type server struct {
	DB       *db.Queries
	Redis    *redis.Client
	Logger   *log.Logger
	DevMode  bool
}

func NewServer(db *db.Queries, r *redis.Client, logger *log.Logger, devMode bool) server {
	return server{
		DB:       db,
		Redis:    r,
		Logger:   logger,
		DevMode:  devMode,
	}
}

func (s server) SetupRoutes() {
	http.HandleFunc("/register", s.handleErrors(s.handleRegister))
	http.HandleFunc("/profile", s.handleErrors(s.handleProfile))
	http.HandleFunc("/login", s.handleErrors(s.handleLogin))
	http.HandleFunc("/logout", s.handleErrors(s.handleLogout))

	if s.DevMode {
		fs := http.FileServer(http.Dir("./static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}
}
