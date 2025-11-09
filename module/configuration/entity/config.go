package entity

import (
	"database/sql"
	"time"
)

type ConfigEntity struct {
	DB *sql.DB
}

type ConfigRequest struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name" binding:"required"`
	ConfigValues any       `json:"config_values" binding:"required"`
	Version      uint16    `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
}

type ConfigResponse struct {
	Id           uint      `json:"id"`
	Name         string    `json:"-"`
	ConfigValues any       `json:"config_values" binding:"required"`
	Version      uint16    `json:"version"`
	CreatedAt    time.Time `json:"created_at" binding:"datetime=2025-11-08T15:38:41+07:00"`
}

type GetConfigRequest struct {
	Name    string `form:"name" binding:"required"`
	Version uint16 `json:"version,omitempty"`
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
