package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

/*
This is a migration file to migrate the schema to SQLite database
This file is needed because at the moment of this program was created
golang migrate doesn't support migration for SQLite yet
*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./db/configs.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create sqlite3 instance: %v", err)
	}

	fSrc, err := (&file.File{}).Open("./db/migrations")
	if err != nil {
		log.Fatalf("Failed to open migration files: %v", err)
	}

	m, err := migrate.NewWithInstance(
		"file",
		fSrc,
		"sqlite3",
		instance,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Migration up completed successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Migration down completed successfully.")
	default:
		log.Fatal("Invalid direction. Use 'up' or 'down'.")
	}
}
