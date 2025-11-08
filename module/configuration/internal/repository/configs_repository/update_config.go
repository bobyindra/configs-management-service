package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) UpdateConfigByConfigName(ctx context.Context, name string, obj *entity.Config) error {
	return nil
}
