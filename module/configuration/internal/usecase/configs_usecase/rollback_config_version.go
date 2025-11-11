package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) RollbackConfigVersionByConfigName(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error) {
	// Check the config version exists or not
	getParams := &entity.GetConfigRequest{
		Name:    params.Name,
		Version: params.Version,
	}
	_, err := u.configsRepo.GetConfigByConfigName(ctx, getParams)
	if err != nil {
		return nil, err
	}

	// Execute Rollback
	configs := &entity.ConfigRequest{
		Name:    params.Name,
		Version: params.Version,
	}

	err = u.configsRepo.RollbackConfigVersionByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
