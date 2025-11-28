package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	cmsConfig "github.com/bobyindra/configs-management-service/module/configuration/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/redis/go-redis/v9"
)

var (
	server    *httptest.Server
	router    *gin.Engine
	authToken string
	db        *sql.DB
	cache     *redis.Client
	mRedis    *miniredis.Miniredis
	dbPath    = "./test.db"
)

func Setup() {
	log.Println("Setting up integration test...")
	BuildRedis()
	BuildDatabase()

	// Start Gin Server
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	schemaPath := filepath.Join(getProjectRoot(), "../schema")
	cmsConfig.CONFIGS_SCHEMA_PATH = schemaPath
	cmsCfg := cmsConfig.CmsConfig{
		Database:          db,
		Redis:             cache,
		Router:            router,
		JWTSecret:         "test-secret",
		JWTExpiryDuration: 86400 * time.Second,
	}
	cmsConfig.RegisterCmsHandler(cmsCfg)
	server = httptest.NewServer(router)

	// Login
	Login()
}

func BuildRedis() {
	// Start Miniredis
	var err error
	mRedis, err = miniredis.Run()
	if err != nil {
		log.Fatalf("Failed to start miniredis: %v", err)
	}

	// Build Redis Client
	cache = redis.NewClient(&redis.Options{
		Addr:         mRedis.Addr(),
		Password:     "",
		DB:           0,
		ReadTimeout:  300 * time.Millisecond,
		WriteTimeout: 300 * time.Millisecond,
	})
}

func BuildDatabase() {
	// Remove existing test DB
	os.Remove(dbPath)

	// Init DB
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed connecting to sqlite DB: %v", err)
	}

	// Run Migration
	migrationPath := filepath.Join(getProjectRoot(), "../db/migration")
	runMigration(migrationPath, db)

	// Seed Database
	query := "INSERT INTO users (username, crypted_password, ROLE) VALUES ('rwuser', '$2a$10$gtMdblzuRU0DU5QyElkSPOC0b6v3XBdFvPRwsQZ98RZSTBoMBKS.C', 'rw')"
	_, _ = db.Exec(query)
}

func Login() {
	var err error
	authToken, err = getAuthToken()
	if err != nil {
		log.Fatalf("Failed to login : %v", err)
	}
}

func TearDown(code int) {
	// TEARDOWN
	log.Println("Cleaning up...")
	server.Close()
	cache.Close()
	mRedis.Close()
	db.Close()
	os.Remove(dbPath)
	os.Exit(code)
}

func TestMain(m *testing.M) {
	Setup()

	// RUN TEST
	code := m.Run()

	TearDown(code)
}

func runMigration(path string, db *sql.DB) {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create sqlite3 instance: %v", err)
	}

	fSrc, err := (&file.File{}).Open(path)
	if err != nil {
		log.Fatalf("Failed to open migration files: %v", err)
	}

	m, err := migrate.NewWithInstance(
		"file",
		fSrc,
		"sqlite3",
		instance,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up failed: %v", err)
	}
}

func getAuthToken() (string, error) {
	body := `{"username":"rwuser","password":"readwriteuser"}`
	resp, err := http.Post(server.URL+"/api/v1/auth/login", "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&result)
	return result.Data.Token, nil
}

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..")
}
