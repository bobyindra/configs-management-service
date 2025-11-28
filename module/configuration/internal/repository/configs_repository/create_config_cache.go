package configs_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsCacheRepository) CreateConfigCache(ctx context.Context, obj *entity.Config) error {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("configs-%s", obj.Name)
	ttl := 12 * time.Hour // Set to 12 hours because configs don't change often
	return r.cache.Set(ctx, key, jsonData, ttl).Err()
}
