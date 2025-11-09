package entity

import (
	"database/sql"
	"time"
)

type ConfigEntity struct {
	DB *sql.DB
}

type ConfigRequest struct {
	Id           int       `json:"-"`
	Name         string    `json:"name" binding:"required"`
	ConfigValues string    `json:"config_values" binding:"required"`
	Version      int       `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

type ConfigResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"-"`
	ConfigValues string    `json:"config_values" binding:"required"`
	Version      int       `json:"version"`
	CreatedAt    time.Time `json:"created_at" binding:"datetime=2025-11-08T15:38:41+07:00"`
}

type GetConfigRequest struct {
	Name    string `form:"name" binding:"required"`
	Version int    `json:"version" binding:"omitempty"`
}

type GetListConfigVersionsRequest struct {
	Name   string `form:"name" binding:"required"`
	Limit  uint32 `form:"limit"`
	Offset uint32 `form:"offset"`
}

func (req *ConfigRequest) ToResponse() *ConfigResponse {
	return &ConfigResponse{
		Id:           req.Id,
		Name:         req.Name,
		ConfigValues: req.ConfigValues,
		Version:      req.Version,
		CreatedAt:    req.CreatedAt,
	}
}
