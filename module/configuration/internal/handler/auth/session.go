package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type SessionHandler struct {
	auth        auth.Auth
	sessionUscs usecase.SessionUsecase
}

func NewSession(auth auth.Auth, sessionUscs usecase.SessionUsecase) *SessionHandler {
	return &SessionHandler{
		auth:        auth,
		sessionUscs: sessionUscs,
	}
}
