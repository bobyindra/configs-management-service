package configshandler

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type ConfigsHandler struct {
	configsUscs usecase.ConfigsManagementUsecase
}

func NewConfigsHandler(configUscs usecase.ConfigsManagementUsecase) *ConfigsHandler {
	return &ConfigsHandler{
		configsUscs: configUscs,
	}
}
