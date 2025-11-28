package usecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/encryption"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"

	authUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/auth"
	configsUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/configs_usecase"
)

type UsecaseList struct {
	AuthUsecase       SessionUsecase
	ConfigsManagement ConfigsManagementUsecase
}

func NewUsecaseList(repoList repository.RepositoryList) UsecaseList {
	return UsecaseList{
		AuthUsecase:       newAuthUsecase(repoList),
		ConfigsManagement: newConfigsManagementUsecase(repoList),
	}
}

func newAuthUsecase(repoList repository.RepositoryList) SessionUsecase {
	encryption := encryption.NewEncryption()
	return authUscs.NewSessionUscs(encryption, repoList.UserRepo)
}

func newConfigsManagementUsecase(repoList repository.RepositoryList) ConfigsManagementUsecase {
	return configsUscs.NewConfigsUsecase(repoList.ConfigsDBRepo, repoList.ConfigsCacheRepo)
}

type ConfigsManagementUsecase interface {
	CreateConfig(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error)
	GetConfigByConfigName(ctx context.Context, params *entity.GetConfigRequest) (*entity.ConfigResponse, error)
	GetListVersionsByConfigName(ctx context.Context, params *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error)
	UpdateConfigByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error)
	RollbackConfigVersionByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error)
}

type SessionUsecase interface {
	Login(ctx context.Context, param *entity.LoginRequest) (*entity.LoginResponse, error)
}
