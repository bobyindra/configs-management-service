package configsusecase

import (
	"context"
	"log"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetConfigByConfigName(ctx context.Context, params *entity.GetConfigRequest) (*entity.ConfigResponse, error) {
	// Retrieve from cache
	cache, err := u.configsRepo.GetConfigCache(ctx, params.Name)
	if err == nil && cache != nil {
		log.Printf("Cache hit for key %s", params.Name)
		return cache.ToResponse(), nil
	}

	// Fallback to database
	configs, err := u.configsRepo.GetConfigByConfigName(ctx, params)
	if err != nil {
		return nil, err
	}

	// Store in cache
	data := &entity.Config{
		Id:           configs.Id,
		Name:         configs.Name,
		ConfigValues: configs.ConfigValues,
		Version:      configs.Version,
		ActorId:      configs.ActorId,
		CreatedAt:    configs.CreatedAt,
	}
	err = u.configsRepo.CreateConfigCache(ctx, data)
	if err != nil {
		log.Printf("Failed to create config cache: %v", err)
	}

	return configs, nil
}
