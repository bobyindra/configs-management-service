package usecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

type ConfigsManagementUsecase interface {
	CreateConfigs(ctx context.Context, obj *entity.Config) error
	GetConfigByConfigName(ctx context.Context, name string) (*entity.Config, error)
	GetListVersionsByConfigName(ctx context.Context, name string) ([]*entity.Config, error)
	UpdateConfigByConfigName(ctx context.Context, name string, obj *entity.Config) error
	RollbackConfigVersionByConfigName(ctx context.Context, name string, version int) error
}
