package configsusecase

import "github.com/bobyindra/configs-management-service/module/configuration/internal/repository"

type configsUsecase struct {
	configsDBRepo    repository.ConfigsManagementDBRepository
	configsCacheRepo repository.ConfigsManagementCacheRepository
}

func NewConfigsUsecase(
	configsDBRepo repository.ConfigsManagementDBRepository,
	configsCacheRepo repository.ConfigsManagementCacheRepository,
) *configsUsecase {
	return &configsUsecase{
		configsDBRepo:    configsDBRepo,
		configsCacheRepo: configsCacheRepo,
	}
}
