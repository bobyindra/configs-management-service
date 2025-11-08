package configshandler

import "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"

type configs struct {
	configsUscs usecase.ConfigsManagementUsecase
}

func NewConfigsHandler(configUscs usecase.ConfigsManagementUsecase) *configs {
	return &configs{
		configsUscs: configUscs,
	}
}
