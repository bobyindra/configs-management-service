package repository

import (
	"context"
	"database/sql"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	userRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
)

type RepositoryList struct {
	ConfigsRepo ConfigsManagementRepository
	UserRepo    UserRepository
}

func NewRepositoryList(db *sql.DB) RepositoryList {
	return RepositoryList{
		ConfigsRepo: NewConfigsManagementRepository(db),
		UserRepo:    NewUserRepository(db),
	}
}

func NewConfigsManagementRepository(db *sql.DB) ConfigsManagementRepository {
	return configsRepository.NewConfigsRepository(db)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository.NewUserRepository(db)
}

type ConfigsManagementRepository interface {
	CreateConfig(ctx context.Context, obj *entity.ConfigRequest) error
	GetConfigByConfigName(ctx context.Context, obj *entity.GetConfigRequest) (*entity.ConfigResponse, error)
	GetListVersionsByConfigName(ctx context.Context, obj *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error)
	UpdateConfigByConfigName(ctx context.Context, obj *entity.ConfigRequest) error
	RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.ConfigRequest) error
}

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}
