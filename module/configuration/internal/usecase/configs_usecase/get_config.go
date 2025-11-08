package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) GetConfigByConfigName(ctx context.Context, name string) (obj *entity.Config, err error) {
	return nil, nil
}
