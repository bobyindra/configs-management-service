package schema

import (
	"os"
	"path/filepath"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

const (
	PAYMENT_CONFIG  = "payment-config"
	FEE_CONFIG      = "fee-config"
	DISCOUNT_CONFIG = "discount-config"
	BANNER_CONFIG   = "banner-config"
	BCA_ENABLED     = "bca-enabled"
	EMAIL_CONFIG    = "email-config"
	WORDING_CONFIG  = "wording-config"
)

var schemaFileMap = map[string]string{
	PAYMENT_CONFIG:  "payment_config.json",
	FEE_CONFIG:      "fee_config.json",
	DISCOUNT_CONFIG: "discount_config.json",
	BANNER_CONFIG:   "banner_config.json",
	BCA_ENABLED:     "bca_enabled.json",
	EMAIL_CONFIG:    "email_config.json",
	WORDING_CONFIG:  "wording_config.json",
}

type schemaRegistry struct {
	basePath string
}

func NewSchemaRegistry(basePath string) *schemaRegistry {
	return &schemaRegistry{
		basePath: basePath,
	}
}

type SchemaRegistry interface {
	GetSchemaByConfigName(configName string) ([]byte, error)
}

func (sr *schemaRegistry) GetSchemaByConfigName(cfgName string) ([]byte, error) {
	fileName, ok := schemaFileMap[cfgName]
	if !ok {
		return nil, entity.ErrConfigSchemaNotFound
	}

	fullPath := filepath.Join(sr.basePath, fileName)
	return os.ReadFile(fullPath)
}
