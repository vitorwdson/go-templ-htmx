package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vitorwdson/go-templ-htmx/db"
	"github.com/vitorwdson/go-templ-htmx/handler"
)

func main() {
	devMode := flag.Bool("dev", false, "Use develoment mode")
	runMigrations := flag.Bool("migrate", false, "Applies migrations and exits the program")
	flag.Parse()

	if *devMode {
		godotenv.Load()
	}

	dbConnection := db.MustConnect()
	defer dbConnection.Close()

	if *runMigrations {
		db.RunMigrations(dbConnection)
		os.Exit(0)
	}

	redis := db.MustConnectRedis()
	logger := log.Default()

	s := handler.NewServer(dbConnection, redis, logger, *devMode)
	s.SetupRoutes()

	logger.Println("Listening on port 3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
