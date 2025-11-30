package configsusecase

import (
	"context"
	"log"
	"time"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) UpdateConfigByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	// Validate schema
	err := u.validateConfigSchema(params.Name, params.ConfigValues)
	if err != nil {
		return nil, err
	}

	// Check the config exists or not
	getParams := &entity.GetConfigRequest{
		Name: params.Name,
	}
	resp, err := u.configsDBRepo.GetConfigByConfigName(ctx, getParams)
	if err != nil {
		return nil, err
	}

	// Validate the values equal with current value or not
	equal, err := util.JsonByteEqual(params.ConfigValues, resp.ConfigValues)
	if err != nil {
		return nil, err
	}
	if equal {
		return nil, entity.ErrNoChangesFound
	}

	// Update the config
	configs := &entity.Config{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
		ActorId:      params.ActorId,
		CreatedAt:    time.Now().UTC(),
	}

	err = u.configsDBRepo.UpdateConfigByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	// Update cache
	err = u.configsCacheRepo.CreateConfigCache(ctx, configs)
	if err != nil {
		log.Printf("Failed to create config cache: %v", err)
	}

	return configs.ToResponse(), nil
}
