package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) RollbackConfigVersionByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	// Check the config version exists or not
	getParams := &entity.GetConfigRequest{
		Name: params.Name,
	}
	resp, err := u.configsRepo.GetConfigByConfigName(ctx, getParams)
	if err != nil {
		return nil, err
	}

	// Validate the version equal with current version or not
	if resp.Version == params.Version {
		return nil, entity.ErrRollbackNotAllowed
	}

	// Execute Rollback
	configs := &entity.Config{
		Name:    params.Name,
		Version: params.Version,
		ActorId: params.ActorId,
	}

	err = u.configsRepo.RollbackConfigVersionByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
