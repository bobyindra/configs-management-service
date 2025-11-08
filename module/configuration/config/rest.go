package config

import (
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configsHandler"
	"github.com/gin-gonic/gin"
)

func RegisterCmsHandler(cfg CmsConfig) error {
	repoList := NewRepositoryList(cfg.Database)
	uscsList := NewUsecaseList(repoList)

	registerConfigsHandler(cfg.Router, uscsList)

	return nil
}

func registerConfigsHandler(router *gin.Engine, uscsList usecaseList) {
	ch := configsHandler.NewConfigsHandler(uscsList.configsManagement)
	v1 := router.Group("/api/v1/configs")
	{
		v1.POST("/:name", ch.CreateConfigs)
		v1.GET("/:name", ch.GetByConfigName)
		v1.GET("/:name/versions", ch.GetListVersionsByConfigName)
		v1.PUT("/:name", ch.UpdateByConfigName)
		v1.POST("/:name/rollback/:version", ch.RollbackVersionByConfigName)
	}
}
