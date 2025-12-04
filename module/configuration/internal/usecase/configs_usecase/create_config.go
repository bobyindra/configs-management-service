package configsusecase

import (
	"context"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) CreateConfig(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	// Validate schema
	err := u.validateConfigSchema(params.Name, params.ConfigValues)
	if err != nil {
		return nil, err
	}

	// Create config entity
	configs := &entity.Config{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
		ActorId:      params.ActorId,
		CreatedAt:    time.Now().UTC(),
	}

	// Store in DB
	err = u.configsDBRepo.CreateConfig(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
