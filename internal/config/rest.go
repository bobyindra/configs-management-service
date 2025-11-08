package config

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	cmsConfig "github.com/bobyindra/configs-management-service/module/configuration/config"
	cmsEntity "github.com/bobyindra/configs-management-service/module/configuration/entity"
)

type RestServer struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

type CmsRestApp struct {
	svcCfg *ServiceConfig

	// Dependencies
	Database *sql.DB
}

func NewCmsRest() (*CmsRestApp, error) {
	app := &CmsRestApp{}
	var err error

	app.svcCfg, err = LoadConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = app.svcCfg.BuildDatabase()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func NewRestServer(app *CmsRestApp) *RestServer {
	r := gin.Default()

	entities := cmsEntity.NewEntities(app.Database)

	cmsCfg := cmsConfig.CmsConfig{
		Database:          entities,
		Router:            r,
		JWTSecret:         app.svcCfg.JWTSecret,
		JWTExpiryDuration: app.svcCfg.JWTExpiryDuration,
	}
	cmsConfig.RegisterCmsHandler(cmsCfg)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", app.svcCfg.RestApiHost, app.svcCfg.RestApiPort),
		Handler: c.Handler(r),
	}

	restServer := &RestServer{
		httpServer:      httpServer,
		shutdownTimeout: time.Duration(app.svcCfg.RestApiShutdownTimeout) * time.Second,
	}

	return restServer
}

func (rs *RestServer) ApiAddress() string {
	return rs.httpServer.Addr
}

func (rs *RestServer) Serve() error {
	return rs.httpServer.ListenAndServe()
}

func (rs *RestServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), rs.shutdownTimeout)
	defer cancel()

	return rs.httpServer.Shutdown(ctx)
}
