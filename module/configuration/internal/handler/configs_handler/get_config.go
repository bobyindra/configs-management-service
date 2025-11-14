package configshandler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/gin-gonic/gin"
)

func (h *ConfigsHandler) GetConfig(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	claim, err := h.auth.ValidateClaim(ctx, r)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	if claim.Role == "rw" || claim.Role == "ro" {
		name := c.Param("name")
		listConfigsParam, err := h.normalizeGetConfigRequest(r.URL.Query())
		if err != nil {
			helper.BuildFailedResponse(w, err)
			return
		}

		listConfigsParam.Name = name

		resp, err := h.configsUscs.GetConfigByConfigName(ctx, listConfigsParam)
		if err != nil {
			helper.BuildFailedResponse(w, err)
			return
		}

		helper.BuildSuccessResponse(w, helper.APIResponse{
			Status: http.StatusOK,
			Data:   resp,
		})
	} else {
		helper.BuildFailedResponse(w, entity.ErrForbidden)
	}
}

func (h *ConfigsHandler) normalizeGetConfigRequest(query url.Values) (*entity.GetConfigRequest, error) {
	param := &entity.GetConfigRequest{}

	if ver := query.Get("version"); ver != "" {
		if v, err := strconv.Atoi(ver); err == nil {
			param.Version = uint16(v)
		} else {
			return nil, entity.WrapError(err)
		}
	}

	return param, nil
}
