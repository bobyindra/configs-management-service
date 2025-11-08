package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) GetConfigByConfigName(ctx context.Context, name string) (obj *entity.Config, err error) {
	return nil, nil
}
