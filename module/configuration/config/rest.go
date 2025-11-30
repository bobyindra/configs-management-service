package config

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	authHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/auth"
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/bobyindra/configs-management-service/module/configuration/schema"
	"github.com/gin-gonic/gin"
)

var CONFIGS_SCHEMA_PATH = "./module/configuration/schema/"

func RegisterCmsHandler(cfg CmsConfig) error {
	// The hardcoded schema registry is temporary and will be moved to database.
	schemaRegistry := schema.NewSchemaRegistry(CONFIGS_SCHEMA_PATH)

	repoList := repository.NewRepositoryList(cfg.Database, cfg.Redis)
	uscsList := usecase.NewUsecaseList(repoList, schemaRegistry)
	authUtil := auth.NewAuth([]byte(cfg.JWTSecret), cfg.JWTExpiryDuration)
	middleware := middleware.NewMiddleware(authUtil)

	registerSessionHandler(cfg.Router, authUtil, uscsList)
	registerConfigsHandler(cfg.Router, middleware, uscsList)

	return nil
}

func registerSessionHandler(router *gin.Engine, auth auth.Auth, uscsList usecase.UsecaseList) {
	sh := authHandler.NewSession(auth, uscsList.AuthUsecase)
	v1 := router.Group("/api/v1/auth")
	{
		v1.POST("/login", sh.Login)
	}
}

func registerConfigsHandler(router *gin.Engine, middleware middleware.MiddlewareInterface, uscsList usecase.UsecaseList) {
	ch := configsHandler.NewConfigsHandler(uscsList.ConfigsManagementUsecase)
	v1 := router.Group("/api/v1/configs")
	v1.Use(middleware.ValidateSession)
	{
		v1.POST("/:name", ch.CreateConfigs)
		v1.GET("/:name", ch.GetConfig)
		v1.GET("/:name/versions", ch.GetConfigVersions)
		v1.PUT("/:name", ch.UpdateConfig)
		v1.POST("/:name/rollback", ch.RollbackConfigVersion)
	}
}
