package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	dbConnection := db.Connect()
	defer dbConnection.Close()

	if *runMigrations {
		db.RunMigrations(dbConnection)
		os.Exit(0)
	}

	redis := db.ConnectRedis()

	app := echo.New()

	h := handler.Handler{
		DB:    dbConnection,
		Redis: redis,
	}
	h.SetupRoutes(app)

	if *devMode {
		fs := http.FileServer(http.Dir("./static/"))
		app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	}

	app.Start(":3333")
}
