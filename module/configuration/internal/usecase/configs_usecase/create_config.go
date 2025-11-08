package configsusecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) CreateConfig(ctx context.Context, name string, obj *entity.Config) error {
	return nil
}
