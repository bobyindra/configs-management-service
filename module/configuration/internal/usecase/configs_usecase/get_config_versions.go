package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetListVersionsByConfigName(ctx context.Context, params *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error) {
	configs, pagination, err := u.configsDBRepo.GetListVersionsByConfigName(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return configs, pagination, nil
}
