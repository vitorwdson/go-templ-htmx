package main

import (
	"github.com/joho/godotenv"
	"github.com/vitorwdson/go-templ-htmx/db"
)

func main() {
	godotenv.Load()

	dbConnection := db.MustConnect()
	defer dbConnection.Close()

	db.RunMigrations(dbConnection)
}
