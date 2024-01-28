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

func parseFlags() (devMode bool, runMigrations bool) {
	devModeFlag := flag.Bool(
		"dev",
		false,
		"Use develoment mode",
	)

	runMigrationsFlag := flag.Bool(
		"migrate",
		false,
		"Applies migrations and exits the program",
	)

	flag.Parse()
	return *devModeFlag, *runMigrationsFlag
}

func main() {
	devMode, runMigrations := parseFlags()

	if devMode {
		godotenv.Load()
	}

	dbConnection := db.MustConnect()
	defer dbConnection.Close()

	if runMigrations {
		db.RunMigrations(dbConnection)
		os.Exit(0)
	}

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
