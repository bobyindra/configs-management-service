package usecase

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"

	configsUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/configs_usecase"
)

type UsecaseList struct {
	ConfigsManagement ConfigsManagementUsecase
}

func NewUsecaseList(repoList repository.RepositoryList) UsecaseList {
	return UsecaseList{
		ConfigsManagement: NewConfigsManagementUsecase(repoList),
	}
}

func NewConfigsManagementUsecase(repoList repository.RepositoryList) ConfigsManagementUsecase {
	return configsUscs.NewConfigsUsecase(repoList.ConfigsRepo)
}

type ConfigsManagementUsecase interface {
	CreateConfig(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error)
	GetConfigByConfigName(ctx context.Context, params *entity.GetConfigRequest) (*entity.ConfigResponse, error)
	GetListVersionsByConfigName(ctx context.Context, params *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, error)
	UpdateConfigByConfigName(ctx context.Context, params *entity.ConfigRequest) (*entity.ConfigResponse, error)
	RollbackConfigVersionByConfigName(ctx context.Context, name string, version int16) (*entity.ConfigResponse, error)
}
