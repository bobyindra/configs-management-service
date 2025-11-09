package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetListVersionsByConfigName(ctx context.Context, params *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, error) {
	return nil, nil
}
