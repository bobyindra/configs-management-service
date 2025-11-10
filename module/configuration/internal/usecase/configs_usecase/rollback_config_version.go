package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) RollbackConfigVersionByConfigName(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error) {
	configs := &entity.ConfigRequest{
		Name:    params.Name,
		Version: params.Version,
	}

	err := u.configsRepo.RollbackConfigVersionByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
