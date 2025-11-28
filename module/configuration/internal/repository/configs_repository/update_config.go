package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var UpdateConfigQuery = "INSERT INTO configs (name, config_values, version, created_at, actor_id) VALUES ($1, $2, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $3, $4) RETURNING id, version"

func (r *configsDBRepository) UpdateConfigByConfigName(ctx context.Context, obj *entity.Config) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// convert the data to json string
	jsonString, err := util.ConvertAnyValueToJsonString(obj.ConfigValues)
	if err != nil {
		return err
	}

	return r.db.QueryRowContext(ctx, UpdateConfigQuery, obj.Name, jsonString, obj.CreatedAt, obj.ActorId).Scan(&obj.Id, &obj.Version)
}
