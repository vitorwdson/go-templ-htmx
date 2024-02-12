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
	http.HandleFunc("GET /register", s.handleErrors(s.handleRegisterGET))
	http.HandleFunc("POST /register", s.handleErrors(s.handleRegisterPOST))
	http.HandleFunc("GET /login", s.handleErrors(s.handleLoginGET))
	http.HandleFunc("POST /login", s.handleErrors(s.handleLoginPOST))
	http.HandleFunc("/logout", s.handleErrors(s.handleLogout))
	http.HandleFunc("/profile", s.handleErrors(s.handleProfile))

	if s.DevMode {
		fs := http.FileServer(http.Dir("./static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}
}
