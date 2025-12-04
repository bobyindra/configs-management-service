package configsusecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *configsUsecase) RollbackConfigVersionByConfigName(ctx context.Context, params *entity.Config) (*entity.ConfigResponse, error) {
	// Check the config version exists or not
	getParams := &entity.GetConfigRequest{
		Name: params.Name,
	}
	resp, err := u.configsDBRepo.GetConfigByConfigName(ctx, getParams)
	if err != nil {
		return nil, err
	}

	// Validate request version bigger than latest version
	if resp.Version < params.Version {
		return nil, entity.ErrConfigVersionNotFound(fmt.Sprint(params.Version))
	}

	// Validate the version equal with current version or not
	if resp.Version == params.Version {
		return nil, entity.ErrRollbackNotAllowed
	}

	// Execute Rollback
	configs := &entity.Config{
		Name:      params.Name,
		Version:   params.Version,
		ActorId:   params.ActorId,
		CreatedAt: time.Now().UTC(),
	}

	err = u.configsDBRepo.RollbackConfigVersionByConfigName(ctx, configs)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	err = u.configsCacheRepo.DeleteConfigCache(ctx, configs.Name)
	if err != nil {
		log.Printf("Failed to delete config cache: %v", err)
	}

	return configs.ToResponse(), nil
}
