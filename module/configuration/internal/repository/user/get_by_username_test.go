package user_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
)

func (s *userRepoSuite) TestUser_GetByUsername_Success() {
	s.Run("Test Get User by username return success", func() {
		// Given
		ctx := context.TODO()
		userData := test.BuildUserData()

		rows := sqlmock.NewRows(user.UserColumns)
		rows.AddRow(
			userData.Id,
			userData.Username,
			userData.CryptedPassword,
			userData.Role,
			userData.CreatedAt,
			userData.UpdatedAt,
		)

		s.mock.ExpectQuery(regexp.QuoteMeta(user.GetUserByUsernameQuery)).WithArgs(userData.Username).WillReturnRows(rows)

		// When
		result, err := s.subject.GetByUsername(ctx, userData.Username)

		// Then
		s.Nil(err)
		s.Equal(userData, result, "Data should be equal")
	})
}

func (s *userRepoSuite) TestUser_GetByUsername_ErrNotFound() {
	s.Run("Test Get User by username return error not found", func() {
		// Given
		ctx := context.TODO()
		username := "test"

		s.mock.ExpectQuery(regexp.QuoteMeta(user.GetUserByUsernameQuery)).WithArgs(username).WillReturnError(sql.ErrNoRows)

		// When
		result, err := s.subject.GetByUsername(ctx, username)

		// Then
		s.Equal(entity.ErrInvalidLogin, err, "Should return ErrNotFound")
		s.Nil(result)
	})
}

func (s *userRepoSuite) TestUser_GetByUsername_ErrDB() {
	s.Run("Test Get User by username return error db", func() {
		// Given
		ctx := context.TODO()
		username := "test"
		mockErr := testutil.ErrUnexpected

		s.mock.ExpectQuery(regexp.QuoteMeta(user.GetUserByUsernameQuery)).WithArgs(username).WillReturnError(mockErr)

		// When
		result, err := s.subject.GetByUsername(ctx, username)

		// Then
		s.EqualError(mockErr, err.Error(), "Should return Error")
		s.Nil(result)
	})
}
