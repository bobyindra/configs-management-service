package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/gin-gonic/gin"
)

func (h *SessionHandler) Login(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	var param entity.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	loginParam, err := h.normalizeLoginRequest(param)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	resp, err := h.sessionUscs.Login(ctx, loginParam)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		return
	}

	resp.Token, err = h.auth.GenerateToken(&auth.TokenParam{Subject: fmt.Sprint(resp.UserID)}, &auth.AdditionalClaim{UserID: resp.UserID, Role: resp.Role})
	if err != nil {
		helper.BuildFailedResponse(w, entity.WrapError(err))
		return
	}

	helper.BuildSuccessResponse(w, helper.APIResponse{
		Status: http.StatusOK,
		Data:   resp,
	})
}

func (h *SessionHandler) normalizeLoginRequest(param entity.LoginRequest) (*entity.LoginRequest, error) {
	if param.Username == "" {
		return nil, entity.ErrEmptyField("username")
	}
	if param.Password == "" {
		return nil, entity.ErrEmptyField("password")
	}

	return util.GeneralNullable(param), nil
}
