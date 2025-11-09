package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetConfigByConfigName(ctx context.Context, params *entity.GetConfigRequest) (*entity.ConfigResponse, error) {
	configs, err := u.configsRepo.GetConfigByConfigName(ctx, params)
	if err != nil {
		return nil, err
	}

	return configs, nil
}
