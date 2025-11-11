package configs_repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/mattn/go-sqlite3"
)

var CreateConfigQuery = "INSERT INTO configs (name, config_values, version, created_at, actor_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"

func (r *configsRepository) CreateConfig(ctx context.Context, obj *entity.Config) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now
	obj.Version = 1

	// convert the data to json string
	jsonData, err := json.Marshal(obj.ConfigValues)
	if err != nil {
		return err
	}
	jsonString := string(jsonData)

	err = r.db.QueryRowContext(ctx, CreateConfigQuery, obj.Name, jsonString, obj.Version, obj.CreatedAt, obj.ActorId).Scan(&obj.Id)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqlite3.ErrNo(sqliteErr.ExtendedCode) == sqlite3.ErrNo(sqlite3.ErrConstraintUnique) {
				return entity.ErrConfigAlreadyExists
			}
		}
		return err
	}

	return nil
}
