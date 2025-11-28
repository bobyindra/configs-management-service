package configs_repository_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/repository"
	configRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type configsRecordSuite struct {
	suite.Suite
}

type configsDBRepoSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject repository.ConfigsManagementDBRepository
	sqlMock sqlmock.Sqlmock
}

type configsCacheRepoSuite struct {
	suite.Suite

	ctrl      *gomock.Controller
	subject   repository.ConfigsManagementCacheRepository
	cacheMock redismock.ClientMock
}

func (s *configsDBRepoSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()

	mocks, repository := test.NewMockConfigsDBRepository(s.ctrl)

	s.subject = repository
	s.sqlMock = mocks.DB
}

func (s *configsCacheRepoSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()

	mocks, repository := test.NewMockConfigsCacheRepository(s.ctrl)

	s.subject = repository
	s.cacheMock = mocks.Cache
}

func (s *configsDBRepoSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsRecordSuite(t *testing.T) {
	suite.Run(t, new(configsRecordSuite))
}

func TestConfigsDBRepository(t *testing.T) {
	suite.Run(t, new(configsDBRepoSuite))
}

func TestConfigsCacheRepository(t *testing.T) {
	suite.Run(t, new(configsCacheRepoSuite))
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

		expected := &entity.ConfigResponse{
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

		expected := &entity.ConfigResponse{
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

func BuildConfigResponseRows(data *entity.Config, rowAdd uint8) *sqlmock.Rows {
	rows := sqlmock.NewRows(configsRepo.ConfigsRepositoryColumns)
	if rowAdd == 0 {
		rowAdd = 1
	}
	for range rowAdd {
		rows.AddRow(
			data.Id,
			data.Name,
			data.ConfigValues,
			data.Version,
			data.CreatedAt,
			data.ActorId,
		)
	}
	println("Build Rows: ", fmt.Sprint(data.CreatedAt))

	return rows
}

func ConfigEntityToConfigResponse(data *entity.Config) *entity.ConfigResponse {
	return &entity.ConfigResponse{
		Id:           data.Id,
		Name:         data.Name,
		ConfigValues: data.ConfigValues,
		Version:      data.Version,
		CreatedAt:    data.CreatedAt,
		ActorId:      data.ActorId,
	}
}
