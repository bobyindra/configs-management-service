package configshandler

import (
	"encoding/json"
	"net/http"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
	"github.com/gin-gonic/gin"
)

func (h *configs) UpdateConfig(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	claim, err := h.auth.ValidateClaim(ctx, r)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}
	if claim.Role != "rw" {
		util.BuildFailedResponse(w, entity.ErrForbidden)
		return
	}

	name := c.Param("name")
	var param entity.ConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		util.BuildFailedResponse(w, err)
		return
	}
	param.Name = name
	param.ActorId = claim.UserID

	// check config name
	createConfigParam, err := h.normalizeCreateConfigRequest(param)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	err = h.validateConfigSchema(param)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	resp, err := h.configsUscs.UpdateConfigByConfigName(ctx, createConfigParam)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	util.BuildSuccessResponse(w, util.APIResponse{
		Status: http.StatusOK,
		Data:   resp,
	})
}
