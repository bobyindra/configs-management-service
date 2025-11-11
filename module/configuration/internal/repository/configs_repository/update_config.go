package configs_repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) UpdateConfigByConfigName(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now

	// convert the data to json string
	jsonData, err := json.Marshal(obj.ConfigValues)
	if err != nil {
		return err
	}
	jsonString := string(jsonData)

	query := "INSERT INTO configs (name, config_values, version, created_at, actor_id) VALUES ($1, $2, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $3, $4) RETURNING id, version"
	return r.db.QueryRowContext(ctx, query, obj.Name, jsonString, obj.CreatedAt, obj.ActorId).Scan(&obj.Id, &obj.Version)
}
