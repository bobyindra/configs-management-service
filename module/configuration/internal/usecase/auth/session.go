package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
)

type sessionUscs struct {
	encryption util.Encryption
	userRepo   repository.UserRepository
}

func NewSessionUscs(encryption util.Encryption, userRepo repository.UserRepository) *sessionUscs {
	return &sessionUscs{
		encryption: encryption,
		userRepo:   userRepo,
	}
}
