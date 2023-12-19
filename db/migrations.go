package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func RunMigrations(db *sql.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            timestamp TIMESTAMP NOT NULL,
            name varchar(60) NOT NULL UNIQUE
        );
    `)
	if err != nil {
		panic(err)
	}

	migrationFiles, err := filepath.Glob("./db/migrations/*.sql")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT name FROM migrations;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	appliedMigrations := map[string]bool{}
	for rows.Next() {
		var m string
		rows.Scan(&m)
		appliedMigrations[m] = true
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sort.Strings(migrationFiles)
	for _, migrationPath := range migrationFiles {
		migrationName := filepath.Base(migrationPath)

		if _, exists := appliedMigrations[migrationName]; exists {
			continue
		}

		file, err := os.ReadFile(migrationPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Running migration %s\n", migrationName)
		migrationSQL := string(file)
		_, err = tx.Exec(migrationSQL)
		if err != nil {
			panic(err)
		}

		_, err = tx.Exec(
			`
			    INSERT INTO migrations (timestamp, name)
			    VALUES ($1, $2);
			`,
			time.Now(),
			migrationName,
		)
		if err != nil {
			panic(err)
		}
	}

	tx.Commit()
}
