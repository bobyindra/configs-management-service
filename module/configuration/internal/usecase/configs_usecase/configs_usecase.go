package configsusecase

import "github.com/bobyindra/configs-management-service/module/configuration/internal/repository"

type configsUsecase struct {
	configsRepo repository.ConfigsManagementRepository
}

func NewConfigsUsecase(configsRepo repository.ConfigsManagementRepository) *configsUsecase {
	return &configsUsecase{
		configsRepo: configsRepo,
	}
}
