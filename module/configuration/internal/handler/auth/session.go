package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type session struct {
	auth        auth.Auth
	sessionUscs usecase.SessionUsecase
}

func NewSession(auth auth.Auth, sessionUscs usecase.SessionUsecase) *session {
	return &session{
		auth:        auth,
		sessionUscs: sessionUscs,
	}
}
