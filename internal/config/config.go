package config

import (
	"database/sql"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/subosito/gotenv"
	"gomodules.xyz/envconfig"
)

type ServiceConfig struct {
	// Database
	DBPath string `envconfig:"DB_PATH" default:"./db/configs.db"`

	// Auth
	JWTSecret         string        `envconfig:"JWT_SECRET"`
	JWTExpiryDuration time.Duration `envconfig:"JWT_EXPIRY_DURATION" default:"86400s"`

	// Redis
	RedisHost         string        `envconfig:"REDIS_HOST" default:"127.0.0.1:6379"`
	RedisPassword     string        `envconfig:"REDIS_PASSWORD"`
	RedisReadTimeout  time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"300ms"`
	RedisWriteTimeout time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"300ms"`

	// Rest
	RestApiHost                      string        `envconfig:"REST_API_HOST" default:"0.0.0.0"`
	RestApiPort                      string        `envconfig:"REST_API_PORT" default:"8080"`
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

func (cfg *ServiceConfig) BuildRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cfg.RedisHost,
		Password:     cfg.RedisPassword,
		DB:           0,
		ReadTimeout:  cfg.RedisReadTimeout,
		WriteTimeout: cfg.RedisWriteTimeout,
	})
}

func (cfg *ServiceConfig) BuildDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
