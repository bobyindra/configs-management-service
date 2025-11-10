package configs_repository

import (
	"database/sql"
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

type configsRepository struct{ db *sql.DB }

func NewConfigsRepository(db *sql.DB) *configsRepository {
	return &configsRepository{
		db: db,
	}
}
