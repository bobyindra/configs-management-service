package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) CreateConfig(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	configs := &entity.Config{
		Name:         params.Name,
		ConfigValues: params.ConfigValues,
		ActorId:      params.ActorId,
	}

	err := u.configsRepo.CreateConfig(ctx, configs)
	if err != nil {
		return nil, err
	}

	return configs.ToResponse(), nil
}
