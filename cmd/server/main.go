package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vitorwdson/go-templ-htmx/db"
	"github.com/vitorwdson/go-templ-htmx/handler"
)

func main() {
	devMode := os.Getenv("DEV_MODE") == "true"

	if devMode {
		godotenv.Load()
	}

	dbConnection := db.MustConnect()
	defer dbConnection.Close()

	redis := db.MustConnectRedis()
	defer redis.Close()

	logger := log.Default()
	s := handler.NewServer(dbConnection, redis, logger, devMode)
	s.SetupRoutes()

	logger.Println("Listening on port 3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
