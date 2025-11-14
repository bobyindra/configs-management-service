package auth

import (
	"github.com/bobyindra/configs-management-service/module/configuration/internal/encryption"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
)

type sessionUscs struct {
	encryption encryption.Encryption
	userRepo   repository.UserRepository
}

func NewSessionUscs(encryption encryption.Encryption, userRepo repository.UserRepository) *sessionUscs {
	return &sessionUscs{
		encryption: encryption,
		userRepo:   userRepo,
	}
}
