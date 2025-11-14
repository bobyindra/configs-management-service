package encryption

import (
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	DefaultBcryptCost = 10
)

type encryption struct{}

func NewEncryption() *encryption {
	return &encryption{}
}

//go:generate mockgen -destination=mock/encryption.go -package=mock --build_flags=--mod=mod github.com/creativeqode/garnet-backend/module/garnet/internal/encryption Encryption
type Encryption interface {
	GeneratePassword(pass string) ([]byte, error)
	ComparePassword(pass1 string, pass2 string) error
}

func (e *encryption) GeneratePassword(pass string) ([]byte, error) {
	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), DefaultBcryptCost)
	if err != nil {
		return nil, entity.WrapError(err)
	}

	return generatedPassword, nil
}

func (e *encryption) ComparePassword(hashedPassword string, plaintext string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintext))
	if err != nil {
		return entity.ErrInvalidLogin
	}

	return nil
}
