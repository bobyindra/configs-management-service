package configs_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) GetConfigByConfigName(ctx context.Context, obj *entity.GetConfigRequest) (*entity.ConfigResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var args []any

	cfgRes := entity.ConfigResponse{}

	query := fmt.Sprintf("SELECT %s FROM configs WHERE name = ?", strings.Join(configsRepositoryColumns, ", "))
	args = append(args, obj.Name)

	if obj.Version > 0 {
		query += " AND version = ?"
		args = append(args, obj.Version)
	} else {
		query += " ORDER BY version DESC LIMIT 1"
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	row := r.db.QueryRowContext(ctx, query, args...)

	err := row.Scan(&cfgRes.Id, &cfgRes.Name, &cfgRes.ConfigValues, &cfgRes.Version, &cfgRes.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound(obj.Name)
		}
		return nil, err
	}

	return util.GeneralNullable(cfgRes), nil
}
