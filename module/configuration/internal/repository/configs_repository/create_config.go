package configs_repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/mattn/go-sqlite3"
)

func (r *configsRepository) CreateConfig(ctx context.Context, obj *entity.ConfigRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	now := time.Now().UTC()
	obj.CreatedAt = now
	obj.Version = 1

	jsonData, err := json.Marshal(obj.ConfigValues)
	if err != nil {
		return err
	}
	jsonString := string(jsonData)

	query := "INSERT INTO configs (name, config_values, version, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err = r.db.QueryRowContext(ctx, query, obj.Name, jsonString, obj.Version, obj.CreatedAt).Scan(&obj.Id)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqlite3.ErrNo(sqliteErr.ExtendedCode) == sqlite3.ErrNo(sqlite3.ErrConstraintUnique) {
				return entity.ErrConfigAlreadyExists
			}
		}
		return err
	}

	return nil
}
