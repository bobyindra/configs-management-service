package config

import (
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/gin-gonic/gin"
)

type CmsConfig struct {
	Database entity.Entity
	Router   *gin.Engine

	JWTSecret         string
	JWTExpiryDuration time.Duration
}
