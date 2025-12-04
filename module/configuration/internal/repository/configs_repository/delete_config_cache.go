package configs_repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (r *configsCacheRepository) DeleteConfigCache(ctx context.Context, key string) error {
	key = fmt.Sprintf("configs-%s", key)
	err := r.cache.Del(ctx, key).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
	}
	return err
}
