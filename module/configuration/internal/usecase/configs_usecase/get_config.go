package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetConfigByConfigName(ctx context.Context, name string) (*entity.ConfigResponse, error) {
	return nil, nil
}
