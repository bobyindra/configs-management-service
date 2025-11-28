package configs_repository

import (
	"context"
	"fmt"
)

func (r *configsCacheRepository) DeleteConfigCache(ctx context.Context, key string) error {
	key = fmt.Sprintf("configs-%s", key)
	return r.cache.Del(ctx, key).Err()
}
