package configs_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

var (
	GetListVersionsConfigQuery       = fmt.Sprintf("SELECT %s FROM configs WHERE name = $1 ORDER BY version DESC LIMIT $2 OFFSET $3", strings.Join(ConfigsRepositoryColumns, ", "))
	GetConfigVersionsTotalCountQuery = "SELECT COUNT(*) FROM configs WHERE name = $1"
)

func (r *configsRepository) GetListVersionsByConfigName(ctx context.Context, obj *entity.GetListConfigVersionsRequest) ([]*entity.ConfigResponse, *entity.PaginationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	limit := obj.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	rows, err := r.db.QueryContext(ctx, GetListVersionsConfigQuery, obj.Name, limit, obj.Offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	resp := make([]*entity.ConfigResponse, 0)
	for rows.Next() {
		result := ConfigRecord{}

		err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.ConfigValues,
			&result.Version,
			&result.CreatedAt,
			&result.ActorId,
		)
		if err != nil {
			return nil, nil, entity.WrapError(err)
		}
		result.ConfigValues = util.ParseAny(result.ConfigValues)
		resp = append(resp, result.ToEntity())
	}

	if len(resp) == 0 {
		return nil, nil, entity.ErrNotFound(obj.Name)
	}

	row := r.db.QueryRowContext(ctx, GetConfigVersionsTotalCountQuery, obj.Name)
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

	return resp, pagination, nil
}
