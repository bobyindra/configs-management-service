package configshandler

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/bobyindra/configs-management-service/module/configuration/schema"
)

type ConfigsHandler struct {
	auth           auth.Auth
	configsUscs    usecase.ConfigsManagementUsecase
	schemaRegistry schema.SchemaRegistry
}

func NewConfigsHandler(auth auth.Auth, configUscs usecase.ConfigsManagementUsecase, sr schema.SchemaRegistry) *ConfigsHandler {
	return &ConfigsHandler{
		auth:           auth,
		configsUscs:    configUscs,
		schemaRegistry: sr,
	}
}
