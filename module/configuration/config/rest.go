package config

import (
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterCmsHandler(cfg CmsConfig) error {
	repoList := repository.NewRepositoryList(cfg.Database)
	uscsList := usecase.NewUsecaseList(repoList)

	registerConfigsHandler(cfg.Router, uscsList)

	return nil
}

func registerConfigsHandler(router *gin.Engine, uscsList usecase.UsecaseList) {
	ch := configsHandler.NewConfigsHandler(uscsList.ConfigsManagement)
	v1 := router.Group("/api/v1/configs")
	{
		v1.POST("/:name", ch.CreateConfigs)
		v1.GET("/:name", ch.GetConfig)
		v1.GET("/:name/versions", ch.GetConfigVersions)
		v1.PUT("/:name", ch.UpdateConfig)
		v1.POST("/:name/rollback/:version", ch.RollbackConfigVersion)
	}
}
