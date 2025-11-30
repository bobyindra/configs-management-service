package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var RollbackConfigVersionQuery = `WITH latest AS (
    SELECT COALESCE(MAX(version), 0) + 1 AS new_version
    FROM configs
    WHERE name = $1
),
src AS (
    SELECT name, config_values
    FROM configs
    WHERE name = $1 AND version = $2
)
INSERT INTO configs (name, config_values, version, created_at, actor_id)
SELECT src.name, src.config_values, latest.new_version, $3, $4
FROM src, latest
RETURNING id, config_values, version`

func (r *configsDBRepository) RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.Config) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := r.db.QueryRowContext(ctx, RollbackConfigVersionQuery, obj.Name, obj.Version, obj.CreatedAt, obj.ActorId).Scan(&obj.Id, &obj.ConfigValues, &obj.Version)
	obj.ConfigValues = util.ParseAny(obj.ConfigValues)
	return err
}
