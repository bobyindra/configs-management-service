package entity

import (
	"database/sql"
	"time"
)

type ConfigEntity struct {
	DB *sql.DB
}

type Config struct {
	Id           int       `json:"id"`
	Name         string    `json:"-"`
	ConfigValues string    `json:"config_values" binding:"required"`
	Version      int       `json:"version"`
	CreatedAt    time.Time `json:"created_at" binding:"datetime=2025-11-08T15:38:41+07:00"`
}
