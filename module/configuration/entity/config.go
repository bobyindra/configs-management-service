package entity

import (
	"database/sql"
	"time"
)

type ConfigEntity struct {
	DB *sql.DB
}

type Config struct {
	Id           uint
	Name         string `form:"name" binding:"required"`
	ConfigValues any    `json:"config_values" binding:"required"`
	Version      uint16 `json:"version"`
	CreatedAt    time.Time
	ActorId      uint
}

type ConfigResponse struct {
	Id           uint      `json:"id"`
	Name         string    `json:"-"`
	ConfigValues any       `json:"config_values" binding:"required"`
	Version      uint16    `json:"version"`
	CreatedAt    time.Time `json:"created_at" binding:"datetime=2025-11-08T15:38:41+07:00"`
	ActorId      uint      `json:"actor_id"`
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

func (req *Config) ToResponse() *ConfigResponse {
	return &ConfigResponse{
		Id:           req.Id,
		Name:         req.Name,
		ConfigValues: req.ConfigValues,
		Version:      req.Version,
		CreatedAt:    req.CreatedAt,
		ActorId:      req.ActorId,
	}
}
