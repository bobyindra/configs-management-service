package configs_repository

import (
	"context"
	"encoding/json"
	"log"
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

	queryGetVersion := "SELECT version FROM configs WHERE name=$1 ORDER BY version DESC"
	err = r.db.QueryRowContext(ctx, queryGetVersion, obj.Name).Scan(&obj.Version)
	if err != nil {
		return err
	}
	obj.Version++
	log.Println(obj.Version)

	query := "INSERT INTO configs (name, config_values, version, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.db.QueryRowContext(ctx, query, obj.Name, jsonString, obj.Version, obj.CreatedAt).Scan(&obj.Id)
}
