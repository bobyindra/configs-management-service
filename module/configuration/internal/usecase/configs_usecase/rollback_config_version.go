package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) RollbackConfigVersionByConfigName(ctx context.Context, name string, version int16) (*entity.ConfigResponse, error) {
	return nil, nil
}
