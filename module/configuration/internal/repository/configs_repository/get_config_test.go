package configs_repository_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/jmoiron/sqlx"
)

func (s *configsDBRepoSuite) TestConfigs_GetConfig_Success() {
	s.Run("Test Get Config by config-name without spesific version return success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		latestVersion := 5
		configData.Version = uint16(latestVersion)

		params := &entity.GetConfigRequest{
			Name: configData.Name,
		}

		expectedResponse := ConfigEntityToConfigResponse(configData)
		rows := BuildConfigResponseRows(configData, 1)
		query := buildQuery(params)
		query = sqlx.Rebind(sqlx.DOLLAR, query)

		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name).
			WillReturnRows(rows)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Nil(err)
		s.Equal(expectedResponse, result, "Result should be equal")
		s.Equal(uint16(latestVersion), result.Version, "Latest version should be equal")
	})

	s.Run("Test Get Config by config-name with spesific version return success", func() {
		// Given
		ctx := context.TODO()

		configData := test.BuildConfigData()
		params := &entity.GetConfigRequest{
			Name:    configData.Name,
			Version: configData.Version,
		}

		expectedResponse := ConfigEntityToConfigResponse(configData)
		rows := BuildConfigResponseRows(configData, 1)
		query := buildQuery(params)

		query = sqlx.Rebind(sqlx.DOLLAR, query)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name, params.Version).
			WillReturnRows(rows)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Nil(err)
		s.Equal(expectedResponse, result, "Result should be equal")
		s.Equal(configData.Version, result.Version, "Result version should be equal with input")
	})
}

func (s *configsDBRepoSuite) TestConfigs_GetConfig_ErrNotFound() {
	s.Run("Test Get Config by config-name without spesific version return err not found", func() {
		// Given
		ctx := context.TODO()

		configData := test.BuildConfigData()
		params := &entity.GetConfigRequest{
			Name: configData.Name,
		}

		query := buildQuery(params)

		query = sqlx.Rebind(sqlx.DOLLAR, query)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name).
			WillReturnError(sql.ErrNoRows)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Equal(entity.ErrNotFound(params.Name), err, "Should return ErrNotFound")
		s.Nil(result)
	})

	s.Run("Test Get Config by config-name with spesific version return err not found", func() {
		// Given
		ctx := context.TODO()

		configData := test.BuildConfigData()
		params := &entity.GetConfigRequest{
			Name:    configData.Name,
			Version: configData.Version,
		}

		query := buildQuery(params)

		query = sqlx.Rebind(sqlx.DOLLAR, query)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name, params.Version).
			WillReturnError(sql.ErrNoRows)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Equal(entity.ErrNotFound(params.Name), err, "Should return ErrNotFound")
		s.Nil(result)
	})
}

func (s *configsDBRepoSuite) TestConfigs_GetConfig_ErrDB() {
	s.Run("Test Get Config by config-name without spesific version return err DB", func() {
		// Given
		ctx := context.TODO()

		configData := test.BuildConfigData()
		params := &entity.GetConfigRequest{
			Name: configData.Name,
		}
		mockErr := testutil.ErrUnexpected

		query := buildQuery(params)

		query = sqlx.Rebind(sqlx.DOLLAR, query)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name).
			WillReturnError(mockErr)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
		s.Nil(result)
	})

	s.Run("Test Get Config by config-name with spesific version return err DB", func() {
		// Given
		ctx := context.TODO()

		configData := test.BuildConfigData()
		params := &entity.GetConfigRequest{
			Name:    configData.Name,
			Version: configData.Version,
		}
		mockErr := testutil.ErrUnexpected

		query := buildQuery(params)

		query = sqlx.Rebind(sqlx.DOLLAR, query)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(params.Name, params.Version).
			WillReturnError(mockErr)

		// When
		result, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
		s.Nil(result)
	})
}

func buildQuery(params *entity.GetConfigRequest) string {
	var args []any
	query := configsRepo.GetConfigQuery
	args = append(args, params.Name)

	if params.Version > 0 {
		query += configsRepo.GetConfigSpecifyVersionQuery
		args = append(args, params.Version)
	} else {
		query += configsRepo.GetConfigOrderByVersionQuery
	}

	return query
}
