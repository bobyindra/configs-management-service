package auth_test

import (
	"context"
	"database/sql"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (s *authUsecaseSuite) TestAuthLogin_Success() {
	s.Run("Test Login - Success", func() {
		// Given
		ctx := context.TODO()
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}

		user := &entity.User{
			Id:              1,
			Username:        params.Username,
			CryptedPassword: params.Password,
			Role:            "rw",
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
		}

		// mock
		s.userRepo.EXPECT().GetByUsername(ctx, params.Username).Return(user, nil)
		s.encryption.EXPECT().ComparePassword(user.CryptedPassword, params.Password).Return(nil)

		// When
		resp, err := s.subject.Login(ctx, params)

		// Then
		s.Nil(err, "Error should be nil")
		s.Equal(user.Id, resp.UserID, "User Id should be equal")
		s.Equal(user.Role, resp.Role, "Role should be equal")
	})
}

func (s *authUsecaseSuite) TestAuthLogin_UserNotFound_Error() {
	s.Run("Test Login - User Not Found Error", func() {
		// Given
		ctx := context.TODO()
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}

		// mock
		s.userRepo.EXPECT().GetByUsername(ctx, params.Username).Return(nil, sql.ErrNoRows)

		// When
		resp, err := s.subject.Login(ctx, params)

		// Then
		s.EqualError(sql.ErrNoRows, err.Error(), "Return sql err no rows")
		s.Nil(resp)
	})
}

func (s *authUsecaseSuite) TestAuthLogin_PasswordNotMatch_Error() {
	s.Run("Test Login - Password Not Match Error", func() {
		// Given
		ctx := context.TODO()
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}

		user := &entity.User{
			Id:              1,
			Username:        params.Username,
			CryptedPassword: params.Password,
			Role:            "rw",
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
		}

		// mock
		s.userRepo.EXPECT().GetByUsername(ctx, params.Username).Return(user, nil)
		s.encryption.EXPECT().ComparePassword(user.CryptedPassword, params.Password).Return(entity.ErrInvalidLogin)

		// When
		resp, err := s.subject.Login(ctx, params)

		// Then
		s.Equal(entity.ErrInvalidLogin, err, "Error should be equal")
		s.Nil(resp)
	})
}
