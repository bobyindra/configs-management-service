package configsusecase

import (
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/kaptinlin/jsonschema"
)

// This Schema Validator is used to validate the predefined schema with the input from user.
// Since this validator is only called from configs usecases but with more than one usecases,
// this validator is created inside the configs_usecase to avoid unnecessary complex separation,
// and in a separate class to keep it clean.
// The unit tests have been covered on the Create and Update usecases unit test.
func (u *configsUsecase) validateConfigSchema(cfgName string, cfgValue any) error {
	// Get schema file from registry
	fileSchema, err := u.schemaRegistry.GetSchemaByConfigName(cfgName)
	if err != nil {
		return err
	}

	// Compile schema
	compiler := jsonschema.NewCompiler()
	sch, err := compiler.Compile(fileSchema)
	if err != nil {
		return err
	}

	// Validate
	result := sch.Validate(cfgValue)
	if result.IsValid() {
		return nil
	}
	return entity.ErrInvalidSchema
}
