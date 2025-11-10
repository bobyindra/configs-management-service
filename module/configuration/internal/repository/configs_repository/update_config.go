package configs_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) UpdateConfigByConfigName(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now

	jsonData, err := json.Marshal(obj.ConfigValues)
	if err != nil {
		return err
	}
	jsonString := string(jsonData)

	queryCheckConfig := "SELECT id FROM configs WHERE name=$1"
	err = r.db.QueryRowContext(ctx, queryCheckConfig, obj.Name).Scan(&obj.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ErrNotFound(obj.Name)
		}
		return err
	}

	query := "INSERT INTO configs (name, config_values, version, created_at) VALUES ($1, $2, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $3) RETURNING id, version"
	return r.db.QueryRowContext(ctx, query, obj.Name, jsonString, obj.CreatedAt).Scan(&obj.Id, &obj.Version)
}
