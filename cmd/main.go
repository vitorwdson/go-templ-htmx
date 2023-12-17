package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/handler/user"
)

func main() {
	app := echo.New()

	userHandler := user.UserHandler{}
	app.GET("/register", userHandler.Register)
	app.GET("/login", userHandler.Login)
	app.GET("/profile", userHandler.Profile)

	app.Start(":3333")
}
