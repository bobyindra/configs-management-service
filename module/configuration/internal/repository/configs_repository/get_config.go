package configs_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var (
	GetConfigQuery               = fmt.Sprintf("SELECT %s FROM configs WHERE name = ?", strings.Join(ConfigsRepositoryColumns, ", "))
	GetConfigSpecifyVersionQuery = " AND version = ?"
	GetConfigOrderByVersionQuery = " ORDER BY version DESC LIMIT 1"
)

func (r *configsRepository) GetConfigByConfigName(ctx context.Context, obj *entity.GetConfigRequest) (*entity.ConfigResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var args []any

	query := GetConfigQuery
	args = append(args, obj.Name)

	if obj.Version > 0 {
		query += GetConfigSpecifyVersionQuery
		args = append(args, obj.Version)
	} else {
		query += GetConfigOrderByVersionQuery
	}

	var cfgRes ConfigRecord
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&cfgRes.Id,
		&cfgRes.Name,
		&cfgRes.ConfigValues,
		&cfgRes.Version,
		&cfgRes.CreatedAt,
		&cfgRes.ActorId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound(obj.Name)
		}
		return nil, err
	}
	cfgRes.ConfigValues = util.ParseAny(cfgRes.ConfigValues)

	return util.GeneralNullable(*cfgRes.ToEntity()), nil
}
