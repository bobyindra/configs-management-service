package user_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository/user"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type userRecordSuite struct {
	suite.Suite
}

type userRepoSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject repository.UserRepository
	mock    sqlmock.Sqlmock
}

func (s *userRepoSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()

	mocks, repository := test.NewMockUserRepository(s.ctrl)

	s.subject = repository
	s.mock = mocks.DB
}

func (s *userRepoSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsRecordSuite(t *testing.T) {
	suite.Run(t, new(userRecordSuite))
}

func TestConfigsRepository(t *testing.T) {
	suite.Run(t, new(userRepoSuite))
}

func (s *userRecordSuite) TestUser_UserToEntity_AllDataProvided() {
	s.Run("Test UserRecord ToEntity with All Data Provided return equal data", func() {
		// Given All User Record Data
		data := user.UserRecord{
			Id:              1,
			Username:        "test",
			CryptedPassword: "ahfjsh123",
			Role:            "rw",
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
		}

		expected := &entity.User{
			Id:              data.Id,
			Username:        data.Username,
			CryptedPassword: data.CryptedPassword,
			Role:            data.Role,
			CreatedAt:       data.CreatedAt,
			UpdatedAt:       data.UpdatedAt,
		}

		// When
		result := data.ToEntity()

		// Then
		s.Equal(expected, result)
	})
}

func (s *userRecordSuite) TestUser_UserToEntity_SomeDataProvided() {
	s.Run("Test UserRecord ToEntity with Some Data Provided return equal data", func() {
		// Given All User Record Data
		data := user.UserRecord{
			Id:              1,
			Username:        "test",
			CryptedPassword: "ahfjsh123",
			Role:            "rw",
		}

		expected := &entity.User{
			Id:              data.Id,
			Username:        data.Username,
			CryptedPassword: data.CryptedPassword,
			Role:            data.Role,
			CreatedAt:       data.CreatedAt,
			UpdatedAt:       data.UpdatedAt,
		}

		// When
		result := data.ToEntity()

		// Then
		s.Equal(expected, result)
	})
}
