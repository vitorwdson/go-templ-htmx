package user

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	DB    *sql.DB
	Redis *redis.Client
}

func (h *UserHandler) SetupRoutes(app *echo.Echo) {
	app.GET("/register", h.Register)
	app.POST("/register", h.PostRegister)

	app.GET("/login", h.Login)
	app.POST("/login", h.PostLogin)

	app.GET("/profile", h.Profile)

	app.Any("/logout", h.Logout)
}
