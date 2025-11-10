package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
)

type session struct {
	encryption util.Encryption
	userRepo   repository.UserRepository
}

func NewSession(encryption util.Encryption, userRepo repository.UserRepository) *session {
	return &session{
		encryption: encryption,
		userRepo:   userRepo,
	}
}
