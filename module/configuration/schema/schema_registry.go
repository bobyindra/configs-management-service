package schema

import (
	"os"

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

func GetSchemaByConfigName(cfgName string) ([]byte, error) {
	switch cfgName {
	case PAYMENT_CONFIG:
		return os.ReadFile("./module/configuration/schema/payment_config.json")
	case FEE_CONFIG:
		return os.ReadFile("./module/configuration/schema/fee_config.json")
	case DISCOUNT_CONFIG:
		return os.ReadFile("./module/configuration/schema/discount_config.json")
	case BANNER_CONFIG:
		return os.ReadFile("./module/configuration/schema/banner_config.json")
	case BCA_ENABLED:
		return os.ReadFile("./module/configuration/schema/bca_enabled.json")
	case EMAIL_CONFIG:
		return os.ReadFile("./module/configuration/schema/email_config.json")
	case WORDING_CONFIG:
		return os.ReadFile("./module/configuration/schema/wording_config.json")
	default:
		return nil, entity.ErrConfigSchemaNotFound
	}
}
