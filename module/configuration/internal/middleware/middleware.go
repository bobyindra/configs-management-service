package middleware

import (
	"github.com/bobyindra/configs-management-service/module/configuration/helper"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyAdditionalClaim string = "AdditionalClaim"
)

type Middleware struct {
	auth auth.Auth
}

type MiddlewareInterface interface {
	ValidateSession(c *gin.Context)
}

func NewMiddleware(auth auth.Auth) *Middleware {
	return &Middleware{
		auth: auth,
	}
}

func (m *Middleware) ValidateSession(c *gin.Context) {
	r := c.Request
	w := c.Writer
	ctx := r.Context()

	claim, err := m.auth.ValidateClaim(ctx, r)
	if err != nil {
		helper.BuildFailedResponse(w, err)
		c.Abort()
		return
	}

	c.Set(ContextKeyAdditionalClaim, &claim.AdditionalClaim)
	c.Next()
}
