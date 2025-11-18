package configshandler

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/bobyindra/configs-management-service/module/configuration/schema"
)

type ConfigsHandler struct {
	configsUscs    usecase.ConfigsManagementUsecase
	schemaRegistry schema.SchemaRegistry
}

func NewConfigsHandler(configUscs usecase.ConfigsManagementUsecase, sr schema.SchemaRegistry) *ConfigsHandler {
	return &ConfigsHandler{
		configsUscs:    configUscs,
		schemaRegistry: sr,
	}
}
