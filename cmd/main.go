package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/handler/user"
)

func main() {
	app := echo.New()

	userHandler := user.UserHandler{}
	app.GET("/register", userHandler.Register)
	app.GET("/login", userHandler.Login)
	app.GET("/profile", userHandler.Profile)

	_, devMode := os.LookupEnv("DEVELOPMENT")
	if devMode {
		fs := http.FileServer(http.Dir("./static/"))
		app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	}

	app.Start(":3333")
}
