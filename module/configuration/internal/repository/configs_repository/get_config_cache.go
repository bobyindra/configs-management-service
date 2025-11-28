package configs_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/redis/go-redis/v9"
)

func (r *configsCacheRepository) GetConfigCache(ctx context.Context, key string) (*entity.Config, error) {
	key = fmt.Sprintf("configs-%s", key)
	resp, err := r.cache.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("Cache miss for key: %s", key)
		}
		return nil, err
	}

	var cfgData entity.Config
	err = json.Unmarshal([]byte(resp), &cfgData)
	if err != nil {
		return nil, err
	}

	return &cfgData, nil
}
