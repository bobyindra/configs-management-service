package config

import (
	"os"
	"time"

	"github.com/subosito/gotenv"
	"gomodules.xyz/envconfig"
)

type ServiceConfig struct {
	// Auth
	JWTSecret         string        `envconfig:"JWT_SECRET"`
	JWTExpiryDuration time.Duration `envconfig:"JWT_EXPIRY_DURATION" default:"86400s"`

	// Rest
	RestApiHost                      string        `envconfig:"REST_API_HOST" default:"0.0.0.0"`
	RestApiPort                      int           `envconfig:"REST_API_PORT" default:"8080"`
	RestApiShutdownTimeout           time.Duration `envconfig:"REST_API_SHUTDOWN_TIMEOUT" default:"30s"`
	RestApiAllowedCredentialsOrigins string        `envconfig:"REST_API_ALLOWED_CREDENTIALS_ORIGINS" default:"*.configs.com"`
}

func LoadConfig() (*ServiceConfig, error) {
	var cfg ServiceConfig
	// Load from .env if exists
	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return nil, err
		}
	}

	err := envconfig.Process("", &cfg)
	return &cfg, err
}
