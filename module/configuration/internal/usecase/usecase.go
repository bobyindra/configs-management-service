package usecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/encryption"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/schema"

	authUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/auth"
	configsUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/configs_usecase"
)

type UsecaseList struct {
	AuthUsecase              SessionUsecase
	ConfigsManagementUsecase ConfigsManagementUsecase
}

func NewUsecaseList(repoList repository.RepositoryList, scmReg schema.SchemaRegistry) UsecaseList {
	return UsecaseList{
		AuthUsecase:              newAuthUsecase(repoList),
		ConfigsManagementUsecase: newConfigsManagementUsecase(repoList, scmReg),
	}
}

func newAuthUsecase(repoList repository.RepositoryList) SessionUsecase {
	encryption := encryption.NewEncryption()
	return authUscs.NewSessionUscs(encryption, repoList.UserRepo)
}

func newConfigsManagementUsecase(repoList repository.RepositoryList, scmReg schema.SchemaRegistry) ConfigsManagementUsecase {
	return configsUscs.NewConfigsUsecase(repoList.ConfigsDBRepo, repoList.ConfigsCacheRepo, scmReg)
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
