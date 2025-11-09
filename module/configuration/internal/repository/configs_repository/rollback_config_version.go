package configs_repository

import (
	"context"
)

func (r *configsRepository) RollbackConfigVersionByConfigName(ctx context.Context, name string, version int16) error {
	return nil
}
