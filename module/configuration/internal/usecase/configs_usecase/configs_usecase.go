package configsusecase

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/schema"
)

type configsUsecase struct {
	configsDBRepo    repository.ConfigsManagementDBRepository
	configsCacheRepo repository.ConfigsManagementCacheRepository
	schemaRegistry   schema.SchemaRegistry
}

func NewConfigsUsecase(
	configsDBRepo repository.ConfigsManagementDBRepository,
	configsCacheRepo repository.ConfigsManagementCacheRepository,
	schemaRegistry schema.SchemaRegistry,
) *configsUsecase {
	return &configsUsecase{
		configsDBRepo:    configsDBRepo,
		configsCacheRepo: configsCacheRepo,
		schemaRegistry:   schemaRegistry,
	}
}
