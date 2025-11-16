package schema_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/schema"
	"github.com/stretchr/testify/assert"
)

func TestSchema_GetSchemaByConfigName(t *testing.T) {
	t.Parallel()

	t.Run("Get Registered Schema Return Success", func(t *testing.T) {
		t.Parallel()

		// Given
		schemaPath := filepath.Join(getProjectRoot(), "schema")
		schemaRegistry := schema.NewSchemaRegistry(schemaPath)
		cfgName := "payment-config"

		// When
		resp, err := schemaRegistry.GetSchemaByConfigName(cfgName)

		// Then
		assert.Nil(t, err, "Error should be nil")
		assert.NotNil(t, resp, "Should return response")
	})

	t.Run("Get Unregistered Schema Return Error", func(t *testing.T) {
		t.Parallel()

		// Given
		schemaPath := filepath.Join(getProjectRoot(), "schema")
		schemaRegistry := schema.NewSchemaRegistry(schemaPath)
		cfgName := "no-config"

		// When
		resp, err := schemaRegistry.GetSchemaByConfigName(cfgName)

		// Then
		assert.NotNil(t, err, "Error should be nil")
		assert.Nil(t, resp, "Should return response")
	})
}

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..")
}
