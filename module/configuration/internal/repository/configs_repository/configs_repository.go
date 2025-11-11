package configs_repository

import (
	"database/sql"
	"time"
)

var (
	configsRepositoryColumns = []string{
		"id",
		"name",
		"config_values",
		"version",
		"created_at",
		"actor_id",
	}
)

const (
	defaultLimit = 10
	timeout      = 3 * time.Second
)

type configsRepository struct{ db *sql.DB }

func NewConfigsRepository(db *sql.DB) *configsRepository {
	return &configsRepository{
		db: db,
	}
}
