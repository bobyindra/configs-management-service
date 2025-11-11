package configs_repository

import (
	"context"
	"time"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now

	query := "INSERT INTO configs (name, config_values, version, created_at, actor_id) SELECT name, config_values, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $2, $3 FROM configs WHERE name=$1 AND version=$4 RETURNING id, config_values, version"
	err := r.db.QueryRowContext(ctx, query, obj.Name, obj.CreatedAt, obj.ActorId, obj.Version).Scan(&obj.Id, &obj.ConfigValues, &obj.Version)
	obj.ConfigValues = util.ParseAny(obj.ConfigValues)
	return err
}
