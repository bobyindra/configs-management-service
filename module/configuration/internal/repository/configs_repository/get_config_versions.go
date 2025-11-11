package configs_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (r *configsRepository) GetListVersionsByConfigName(ctx context.Context, obj *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	limit := obj.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	query := fmt.Sprintf("SELECT %s FROM configs WHERE name = $1 ORDER BY version DESC LIMIT $2 OFFSET $3", strings.Join(configsRepositoryColumns, ", "))
	rows, err := r.db.QueryContext(ctx, query, obj.Name, limit, obj.Offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := make([]*entity.ConfigResponse, 0)
	for rows.Next() {
		var cfgRes entity.ConfigResponse

		err := rows.Scan(
			&cfgRes.Id,
			&cfgRes.Name,
			&cfgRes.ConfigValues,
			&cfgRes.Version,
			&cfgRes.CreatedAt,
			&cfgRes.ActorId,
		)
		if err != nil {
			return nil, nil, entity.WrapError(err)
		}
		cfgRes.ConfigValues = util.ParseAny(cfgRes.ConfigValues)
		result = append(result, util.GeneralNullable(cfgRes))
	}

	if len(result) == 0 {
		return nil, nil, entity.ErrNotFound(obj.Name)
	}

	countQuery := "SELECT COUNT(*) FROM configs WHERE name = $1"
	row := r.db.QueryRowContext(ctx, countQuery, obj.Name)
	var total uint32
	err = row.Scan(&total)
	if err != nil {
		return nil, nil, entity.WrapError(err)
	}

	pagination := &entity.PaginationResponse{
		OffsetPagination: &entity.OffsetPagination{
			Limit:  limit,
			Offset: obj.Offset,
			Total:  total,
		},
	}

	return result, pagination, nil
}
