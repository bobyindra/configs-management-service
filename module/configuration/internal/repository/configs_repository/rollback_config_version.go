package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var RollbackConfigVersionQuery = "INSERT INTO configs (name, config_values, version, created_at, actor_id) SELECT name, config_values, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $2, $3 FROM configs WHERE name=$1 AND version=$4 RETURNING id, config_values, version"

func (r *configsDBRepository) RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.Config) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := r.db.QueryRowContext(ctx, RollbackConfigVersionQuery, obj.Name, obj.CreatedAt, obj.ActorId, obj.Version).Scan(&obj.Id, &obj.ConfigValues, &obj.Version)
	obj.ConfigValues = util.ParseAny(obj.ConfigValues)
	return err
}
