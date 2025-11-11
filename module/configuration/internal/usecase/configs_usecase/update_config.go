package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) UpdateConfigByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	// Check the config exists or not
	getParams := &entity.GetConfigRequest{
		Name: params.Name,
	}
	resp, err := u.configsRepo.GetConfigByConfigName(ctx, getParams)
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
	}

	err = u.configsRepo.UpdateConfigByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
