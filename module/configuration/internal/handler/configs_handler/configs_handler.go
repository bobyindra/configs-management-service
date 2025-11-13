package configshandler

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type ConfigsHandler struct {
	auth        auth.Auth
	configsUscs usecase.ConfigsManagementUsecase
}

func NewConfigsHandler(auth auth.Auth, configUscs usecase.ConfigsManagementUsecase) *ConfigsHandler {
	return &ConfigsHandler{
		auth:        auth,
		configsUscs: configUscs,
	}
}
