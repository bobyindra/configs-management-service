package configsusecase

import (
	"context"
	"log"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) CreateConfig(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	configs := &entity.Config{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
		ActorId:      params.ActorId,
		CreatedAt:    time.Now().UTC(),
	}

	err := u.configsRepo.CreateConfig(ctx, configs)
	if err != nil {
		return nil, err
	}

	// Store in cache
	err = u.configsRepo.CreateConfigCache(ctx, configs)
	if err != nil {
		log.Printf("Failed to create config cache: %v", err)
	}

	return configs.ToResponse(), nil
}
