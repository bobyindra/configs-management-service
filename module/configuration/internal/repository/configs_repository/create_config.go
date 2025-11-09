package configs_repository

import (
	"context"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) CreateConfig(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now

	query := "INSERT INTO configs (name, config_values, version, created_at) VALUES ($1, $2, $3, $4)"
	return r.db.QueryRowContext(ctx, query, obj.Name, obj.ConfigValues, 1, obj.CreatedAt).Scan(&obj.Id)
}
