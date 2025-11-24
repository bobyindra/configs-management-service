package configs_repository

import (
	"context"
)

func (r *configsRepository) DeleteConfigCache(ctx context.Context, key string) error {
	return r.cache.Del(ctx, key).Err()
}
