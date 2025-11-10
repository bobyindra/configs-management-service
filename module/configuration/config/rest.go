package config

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	authHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/auth"
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterCmsHandler(cfg CmsConfig) error {
	repoList := repository.NewRepositoryList(cfg.Database)
	uscsList := usecase.NewUsecaseList(repoList)
	authUtil := auth.NewAuth([]byte(cfg.JWTSecret), cfg.JWTExpiryDuration)

	registerSessionHandler(cfg.Router, authUtil, uscsList)
	registerConfigsHandler(cfg.Router, authUtil, uscsList)

	return nil
}

func registerSessionHandler(router *gin.Engine, auth auth.Auth, uscsList usecase.UsecaseList) {
	sh := authHandler.NewSession(auth, uscsList.AuthUsecase)
	v1 := router.Group("/api/v1/auth")
	{
		v1.POST("/login", sh.Login)
	}
}

func registerConfigsHandler(router *gin.Engine, auth auth.Auth, uscsList usecase.UsecaseList) {
	ch := configsHandler.NewConfigsHandler(auth, uscsList.ConfigsManagement)
	v1 := router.Group("/api/v1/configs")
	{
		v1.POST("/:name", ch.CreateConfigs)
		v1.GET("/:name", ch.GetConfig)
		v1.GET("/:name/versions", ch.GetConfigVersions)
		v1.PUT("/:name", ch.UpdateConfig)
		v1.POST("/:name/rollback", ch.RollbackConfigVersion)
	}
}
