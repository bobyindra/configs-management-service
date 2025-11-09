package configshandler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
	"github.com/gin-gonic/gin"
)

func (h *configs) GetConfig(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	// TODO: Check Permission

	name := c.Param("name")

	listUnitsParam, err := h.normalizeConfigRequest(r.URL.Query())
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	listUnitsParam.Name = name

	resp, err := h.configsUscs.GetConfigByConfigName(ctx, listUnitsParam)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	util.BuildSuccessResponse(w, util.APIResponse{
		Status: http.StatusOK,
		Data:   resp,
	})
}

func (h *configs) normalizeConfigRequest(query url.Values) (*entity.GetConfigRequest, error) {
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
