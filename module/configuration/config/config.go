package config

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CmsConfig struct {
	Router *gin.Engine

	JWTSecret         string
	JWTExpiryDuration time.Duration
}
