package configshandler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *ConfigsHandler) GetConfigVersions(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	// Get Claim data from Context
	ctxClaim, _ := c.Get(middleware.ContextKeyAdditionalClaim)
	claim, _ := ctxClaim.(*auth.AdditionalClaim)

	if claim.Role == "rw" || claim.Role == "ro" {
		name := c.Param("name")
		listConfigsParam := h.normalizeGetListConfigRequest(r.URL.Query())
		listConfigsParam.Name = name

		resp, pagination, err := h.configsUscs.GetListVersionsByConfigName(ctx, listConfigsParam)
		if err != nil {
			helper.BuildFailedResponse(w, err)
			return
		}

		helper.BuildSuccessResponse(w, helper.APIResponse{
			Status: http.StatusOK,
			Meta:   pagination,
			Data:   resp,
		})
	} else {
		helper.BuildFailedResponse(w, entity.ErrForbidden)
	}
}

func (h *ConfigsHandler) normalizeGetListConfigRequest(query url.Values) *entity.GetListConfigVersionsRequest {
	param := &entity.GetListConfigVersionsRequest{}

	if l := query.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			param.Limit = uint32(v)
		} else {
			param.Limit = 0
		}
	}

	if o := query.Get("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			param.Offset = uint32(v)
		} else {
			param.Offset = 0
		}
	}

	return param
}
