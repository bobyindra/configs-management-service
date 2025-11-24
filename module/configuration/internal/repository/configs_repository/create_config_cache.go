package configs_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) CreateConfigCache(ctx context.Context, obj *entity.Config) error {
	key := fmt.Sprintf("configs-%s", obj.Name)
	ttl := 12 * time.Hour // Set to 12 hours because configs don't change often
	return r.cache.Set(ctx, key, obj, ttl).Err()
}
