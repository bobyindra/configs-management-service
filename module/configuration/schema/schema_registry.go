package schema

import (
	"os"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

const (
	PAYMENT_CONFIG = "payment-config"
	FEE_CONFIG     = "fee-config"
)

func GetSchemaByConfigName(cfgName string) ([]byte, error) {
	switch cfgName {
	case PAYMENT_CONFIG:
		return os.ReadFile("./module/configuration/schema/payment_config.json")
	case FEE_CONFIG:
		return os.ReadFile("./module/configuration/schema/fee_config.json")
	default:
		return nil, entity.ErrConfigNotFound
	}
}
