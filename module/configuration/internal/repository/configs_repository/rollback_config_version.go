package configs_repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now

	queryCheckConfig := "SELECT id FROM configs WHERE name=$1 AND version=$2"
	err := r.db.QueryRowContext(ctx, queryCheckConfig, obj.Name, obj.Version).Scan(&obj.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ErrConfigVersionNotFound(strconv.Itoa(int(obj.Version)))
		}
		return err
	}

	query := "INSERT INTO configs (name, config_values, version, created_at) SELECT name, config_values, (SELECT COALESCE(MAX(version), 0) + 1 FROM configs WHERE name = $1), $2 FROM configs WHERE name=$1 AND version=$3 RETURNING id, config_values, version"
	return r.db.QueryRowContext(ctx, query, obj.Name, obj.CreatedAt, obj.Version).Scan(&obj.Id, &obj.ConfigValues, &obj.Version)
}
