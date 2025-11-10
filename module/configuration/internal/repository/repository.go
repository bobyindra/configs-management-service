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
	CreateConfig(ctx context.Context, obj *entity.ConfigRequest) error
	GetConfigByConfigName(ctx context.Context, obj *entity.GetConfigRequest) (*entity.ConfigResponse, error)
	GetListVersionsByConfigName(ctx context.Context, obj *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error)
	UpdateConfigByConfigName(ctx context.Context, obj *entity.ConfigRequest) error
	RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.ConfigRequest) error
}
