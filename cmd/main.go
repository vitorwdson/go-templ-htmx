package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/db"
	"github.com/vitorwdson/go-templ-htmx/handler/user"
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

	app := echo.New()

	userHandler := user.UserHandler{}
	app.GET("/register", userHandler.Register)
	app.GET("/login", userHandler.Login)

	if *devMode {
		fs := http.FileServer(http.Dir("./static/"))
		app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	}

	app.Start(":3333")
}
