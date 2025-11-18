package configshandler

import (
	"encoding/json"
	"net/http"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *ConfigsHandler) UpdateConfig(c *gin.Context) {
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

	// check config name
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

	resp, err := h.configsUscs.UpdateConfigByConfigName(ctx, createConfigParam)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	helper.BuildSuccessResponse(w, helper.APIResponse{
		Status: http.StatusOK,
		Data:   resp,
	})
}
