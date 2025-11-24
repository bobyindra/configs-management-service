package config

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type CmsConfig struct {
	Database *sql.DB
	Redis    *redis.Client
	Router   *gin.Engine

	JWTSecret         string
	JWTExpiryDuration time.Duration
}
