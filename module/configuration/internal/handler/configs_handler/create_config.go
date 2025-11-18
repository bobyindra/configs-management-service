package configshandler

import (
	"encoding/json"
	"net/http"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/kaptinlin/jsonschema"
)

func (h *ConfigsHandler) CreateConfigs(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	// Get Claim data from Context
	ctxClaim, _ := c.Get(middleware.ContextKeyAdditionalClaim)
	claim, _ := ctxClaim.(*auth.AdditionalClaim)

	if claim.Role != "rw" {
		helper.BuildFailedResponse(w, entity.ErrForbidden)
		return
	}

	name := c.Param("name")
	var param entity.Config
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}
	param.Name = name
	param.ActorId = claim.UserID

	// normalize params
	createConfigParam, err := h.normalizeCreateConfigRequest(param)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	err = h.validateConfigSchema(param)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	resp, err := h.configsUscs.CreateConfig(ctx, createConfigParam)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	helper.BuildSuccessResponse(w, helper.APIResponse{
		Status: http.StatusCreated,
		Data:   resp,
	})
}

func (h *ConfigsHandler) normalizeCreateConfigRequest(param entity.Config) (*entity.Config, error) {
	if param.Name == "" {
		return nil, entity.ErrEmptyField("name")
	}
	if param.ConfigValues == nil {
		return nil, entity.ErrEmptyField("config_values")
	}

	return &param, nil
}

func (h *ConfigsHandler) validateConfigSchema(param entity.Config) error {
	fileSchema, err := h.schemaRegistry.GetSchemaByConfigName(param.Name)
	if err != nil {
		return err
	}

	compiler := jsonschema.NewCompiler()
	sch, err := compiler.Compile(fileSchema)
	if err != nil {
		return err
	}

	// Validate
	result := sch.Validate(param.ConfigValues)
	if result.IsValid() {
		return nil
	}
	return entity.ErrInvalidSchema
}
