package auth

import (
	"context"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (u *session) Login(ctx context.Context, param *entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := u.userRepo.GetByUsername(ctx, param.Username)
	if err != nil {
		return nil, err
	}

	err = u.encryption.ComparePassword(user.CryptedPassword, param.Password)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		UserID: user.Id,
		Role:   user.Role,
	}, nil
}
