package repository

import (
	"context"
	"database/sql"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	configsRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
)

func NewConfigsManagementRepository(db *sql.DB) repository.ConfigsManagementRepository {
	return configsRepository.NewConfigsRepository(db)
}

type ConfigsManagementRepository interface {
	CreateConfigs(ctx context.Context, obj *entity.Config) error
	GetConfigByConfigName(ctx context.Context, name string) (*entity.Config, error)
	GetListVersionsByConfigName(ctx context.Context, name string) ([]*entity.Config, error)
	UpdateConfigByConfigName(ctx context.Context, name string, obj *entity.Config) error
	RollbackConfigVersionByConfigName(ctx context.Context, name string, version int) error
}
