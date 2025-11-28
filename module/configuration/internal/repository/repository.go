package repository

import (
	"context"
	"database/sql"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	userRepository "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
	"github.com/redis/go-redis/v9"
)

type RepositoryList struct {
	ConfigsDBRepo    ConfigsManagementDBRepository
	ConfigsCacheRepo ConfigsManagementCacheRepository
	UserRepo         UserRepository
}

func NewRepositoryList(db *sql.DB, cache *redis.Client) RepositoryList {
	return RepositoryList{
		ConfigsDBRepo:    newConfigsManagementDBRepository(db),
		ConfigsCacheRepo: newConfigsManagementCacheRepository(cache),
		UserRepo:         newUserRepository(db),
	}
}

func newConfigsManagementDBRepository(db *sql.DB) ConfigsManagementDBRepository {
	return configsRepository.NewConfigsDBRepository(db)
}

func newConfigsManagementCacheRepository(cache *redis.Client) ConfigsManagementCacheRepository {
	return configsRepository.NewConfigsCacheRepository(cache)
}

func newUserRepository(db *sql.DB) UserRepository {
	return userRepository.NewUserRepository(db)
}

type ConfigsManagementDBRepository interface {
	CreateConfig(ctx context.Context, obj *entity.Config) error
	GetConfigByConfigName(ctx context.Context, obj *entity.GetConfigRequest) (*entity.ConfigResponse, error)
	GetListVersionsByConfigName(ctx context.Context, obj *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error)
	UpdateConfigByConfigName(ctx context.Context, obj *entity.Config) error
	RollbackConfigVersionByConfigName(ctx context.Context, obj *entity.Config) error
}

type ConfigsManagementCacheRepository interface {
	CreateConfigCache(ctx context.Context, obj *entity.Config) error
	GetConfigCache(ctx context.Context, key string) (*entity.Config, error)
	DeleteConfigCache(ctx context.Context, key string) error
}

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}
