package entity

import (
	"database/sql"
	"time"
)

type UserEntity struct {
	DB *sql.DB
}

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
