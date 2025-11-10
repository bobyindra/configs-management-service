package configshandler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
	"github.com/gin-gonic/gin"
)

func (h *configs) GetConfigVersions(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	claim, err := h.auth.ValidateClaim(ctx, r)
	if err != nil {
		util.BuildFailedResponse(w, err)
		return
	}

	if claim.Role == "rw" || claim.Role == "ro" {
		name := c.Param("name")
		listConfigsParam, err := h.normalizeGetListConfigRequest(r.URL.Query())
		if err != nil {
			util.BuildFailedResponse(w, err)
			return
		}
		listConfigsParam.Name = name

		resp, pagination, err := h.configsUscs.GetListVersionsByConfigName(ctx, listConfigsParam)
		if err != nil {
			util.BuildFailedResponse(w, err)
			return
		}

		util.BuildSuccessResponse(w, util.APIResponse{
			Status: http.StatusOK,
			Meta:   pagination,
			Data:   resp,
		})
	} else {
		util.BuildFailedResponse(w, entity.ErrForbidden)
	}
}

func (h *configs) normalizeGetListConfigRequest(query url.Values) (*entity.GetListConfigVersionsRequest, error) {
	param := &entity.GetListConfigVersionsRequest{}

	if l := query.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			param.Limit = uint32(v)
		} else {
			return nil, entity.WrapError(err)
		}
	}

	if o := query.Get("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			param.Offset = uint32(v)
		} else {
			return nil, entity.WrapError(err)
		}
	}

	return param, nil
}
