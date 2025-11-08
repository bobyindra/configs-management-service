package repository

import (
	"context"
	"database/sql"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
)

type RepositoryList struct {
	ConfigsRepo ConfigsManagementRepository
}

func NewRepositoryList(db *sql.DB) RepositoryList {
	return RepositoryList{
		ConfigsRepo: NewConfigsManagementRepository(db),
	}
}

func NewConfigsManagementRepository(db *sql.DB) ConfigsManagementRepository {
	return configsRepository.NewConfigsRepository(db)
}

type ConfigsManagementRepository interface {
	CreateConfig(ctx context.Context, name string, obj *entity.Config) error
	GetConfigByConfigName(ctx context.Context, name string) (*entity.Config, error)
	GetListVersionsByConfigName(ctx context.Context, name string) ([]*entity.Config, error)
	UpdateConfigByConfigName(ctx context.Context, name string, obj *entity.Config) error
	RollbackConfigVersionByConfigName(ctx context.Context, name string, version int16) error
}
