package configs_repository

import (
	"database/sql"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/redis/go-redis/v9"
)

var (
	ConfigsRepositoryColumns = []string{
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

type configsRepository struct {
	db    *sql.DB
	cache *redis.Client
}

func NewConfigsRepository(db *sql.DB, cache *redis.Client) *configsRepository {
	return &configsRepository{
		db:    db,
		cache: cache,
	}
}

type ConfigRecord struct {
	Id           uint      `db:"id"`
	Name         string    `db:"name"`
	ConfigValues any       `db:"config_values"`
	Version      uint16    `db:"version"`
	CreatedAt    time.Time `db:"created_at"`
	ActorId      uint      `db:"actor_id"`
}

func (cr ConfigRecord) ToEntity() *entity.ConfigResponse {
	cfgRes := &entity.ConfigResponse{
		Id:           cr.Id,
		Name:         cr.Name,
		ConfigValues: cr.ConfigValues,
		Version:      cr.Version,
		CreatedAt:    cr.CreatedAt,
		ActorId:      cr.ActorId,
	}
	return cfgRes
}
