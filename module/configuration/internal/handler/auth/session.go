package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
)

type session struct {
	auth        auth.Auth
	sessionUscs usecase.Session
}

func NewSession(auth auth.Auth, sessionUscs usecase.Session) *session {
	return &session{
		auth:        auth,
		sessionUscs: sessionUscs,
	}
}
