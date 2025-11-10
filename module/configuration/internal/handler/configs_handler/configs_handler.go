package configshandler

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type configs struct {
	auth        auth.Auth
	configsUscs usecase.ConfigsManagementUsecase
}

func NewConfigsHandler(auth auth.Auth, configUscs usecase.ConfigsManagementUsecase) *configs {
	return &configs{
		auth:        auth,
		configsUscs: configUscs,
	}
}
