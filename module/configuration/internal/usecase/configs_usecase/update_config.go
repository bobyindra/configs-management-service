package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) UpdateConfigByConfigName(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error) {
	configs := &entity.ConfigRequest{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
	}

	err := u.configsRepo.UpdateConfigByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
