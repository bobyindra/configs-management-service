package configs_repository

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) GetConfigCache(ctx context.Context, key string) (*entity.Config, error) {
	var cfgData entity.Config
	err := r.cache.GetSet(ctx, key, &cfgData).Err()
	if err != nil {
		return nil, err
	}

	return &cfgData, nil
}
