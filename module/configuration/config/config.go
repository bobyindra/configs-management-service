package config

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type CmsConfig struct {
	Database *sql.DB
	Router   *gin.Engine

	JWTSecret         string
	JWTExpiryDuration time.Duration
}
