package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) UpdateConfigByConfigName(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error) {
	// Check the config exists or not
	getParams := &entity.GetConfigRequest{
		Name: params.Name,
	}
	_, err := u.configsRepo.GetConfigByConfigName(ctx, getParams)
	if err != nil {
		return nil, err
	}

	// Update the config
	configs := &entity.ConfigRequest{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
	}

	err = u.configsRepo.UpdateConfigByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
