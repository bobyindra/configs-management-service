package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) UpdateConfigByConfigName(ctx context.Context, name string, obj *entity.Config) error {
	return nil
}
