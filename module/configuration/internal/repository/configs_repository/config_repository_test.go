package configs_repository_test

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	configRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type configsRecordSuite struct {
	suite.Suite
}

type configsRepoSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject repository.ConfigsManagementRepository
	mock    sqlmock.Sqlmock
}

func (s *configsRepoSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()

	mocks, repository := test.NewMockConfigsRepository(s.ctrl)

	s.subject = repository
	s.mock = mocks.DB
}

func (s *configsRecordSuite) TestConfigs_ConfigToEntity_AllDataProvided() {
	s.Run("Test ConfigRecord ToEntity with All Data Provided return equal data", func() {
		// Given All User Record Data
		data := configRepo.ConfigRecord{
			Id:           1,
			Name:         "test",
			ConfigValues: "ahfjsh123",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      1,
		}

		expected := &entity.Config{
			Id:           data.Id,
			Name:         data.Name,
			ConfigValues: data.ConfigValues,
			Version:      data.Version,
			CreatedAt:    data.CreatedAt,
			ActorId:      data.ActorId,
		}

		// When
		result := data.ToEntity()

		// Then
		s.Equal(expected, result)
	})
}

func (s *configsRecordSuite) TestConfigs_ConfigToEntity_SomeDataProvided() {
	s.Run("Test ConfigRecord ToEntity with Some Data Provided return equal data", func() {
		// Given All User Record Data
		data := configRepo.ConfigRecord{
			Id:           1,
			Name:         "test",
			ConfigValues: "ahfjsh123",
			Version:      1,
		}

		expected := &entity.Config{
			Id:           data.Id,
			Name:         data.Name,
			ConfigValues: data.ConfigValues,
			Version:      data.Version,
			CreatedAt:    data.CreatedAt,
			ActorId:      data.ActorId,
		}

		// When
		result := data.ToEntity()

		// Then
		s.Equal(expected, result)
	})
}
