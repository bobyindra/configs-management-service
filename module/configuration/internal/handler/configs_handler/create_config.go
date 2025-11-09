package configshandler

import (
	"encoding/json"
	"net/http"

	generalUtil "github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
	"github.com/gin-gonic/gin"
)

func (h *configs) CreateConfigs(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	// TODO: Check Permission

	var param entity.ConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	createConfigParam, err := h.normalizeCreateConfigRequest(param)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	resp, err := h.configsUscs.CreateConfig(ctx, createConfigParam)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	util.BuildSuccessResponse(w, util.APIResponse{
		Status: http.StatusCreated,
		Data:   resp,
	})
}

func (h *configs) normalizeCreateConfigRequest(param entity.ConfigRequest) (*entity.ConfigRequest, error) {
	if param.Name == "" {
		return nil, entity.ErrEmptyField("name")
	}
	if len(param.ConfigValues) == 0 {
		return nil, entity.ErrEmptyField("config_values")
	}

	return generalUtil.GeneralNullable(param), nil
}
