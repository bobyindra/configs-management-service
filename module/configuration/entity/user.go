package entity

import (
	"database/sql"
	"time"
)

type UserEntity struct {
	DB *sql.DB
}

type User struct {
	Id              uint
	Username        string
	CryptedPassword string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID uint   `json:"-"`
	Role   string `json:"-"`
	Token  string `json:"token"`
}
