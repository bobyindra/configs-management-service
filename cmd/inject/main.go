package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

/*
This function will inject the users to the users table
This step is needed for testing purpose because the API for registering an user is not provided yet in this project.
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	now := time.Now().UTC()

	db, err := sql.Open("sqlite3", "./db/configs.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	query := "INSERT INTO users (username, crypted_password, ROLE, created_at, updated_at) VALUES ('rouser', '$2a$10$mAhpmY5gLg5gcefBhQP.fef/ozsOD9zN7.mBvAIwW9CeE9W.1WXpi', 'ro', $1, $1), ('rwuser', '$2a$10$gtMdblzuRU0DU5QyElkSPOC0b6v3XBdFvPRwsQZ98RZSTBoMBKS.C', 'rw', $1, $1)"
	if _, err := db.ExecContext(ctx, query, now); err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqlite3.ErrNo(sqliteErr.ExtendedCode) == sqlite3.ErrNo(sqlite3.ErrConstraintUnique) {
				log.Fatal("Users already exists")
			}
		}
		log.Fatalf("Inject users failed with error: %v", err)
	}
	log.Println("Inject users completed successfully")
}
