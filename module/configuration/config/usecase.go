package config

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	configsUscs "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/configsusecase"
)

type usecaseList struct {
	configsManagement usecase.ConfigsManagementUsecase
}

func NewUsecaseList(repoList repositoryList) usecaseList {
	return configsUscs.NewConfigsUsecase(repoList)
}

func NewConfigsManagementUsecase(repoList repositoryList) usecase.ConfigsManagementUsecase {
	return configsUscs.NewConfigsUsecase(repoList.configsManagement)
}
