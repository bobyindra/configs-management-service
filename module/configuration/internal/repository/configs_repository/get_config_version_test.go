package configs_repository_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
)

func (s *configsRepoSuite) TestConfigs_GetListVersion_Success() {
	s.Run("Test Get Config List Version by config-name return success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		totalRowValue := 2

		params := &entity.GetListConfigVersionsRequest{
			Name:   configData.Name,
			Limit:  1,
			Offset: 0,
		}

		expectedResponse := ConfigEntityToConfigResponse(configData)

		expectedPagination := &entity.PaginationResponse{
			OffsetPagination: &entity.OffsetPagination{
				Limit:  params.Limit,
				Offset: params.Offset,
				Total:  uint32(totalRowValue),
			},
		}

		rows := BuildConfigResponseRows(configData, 2)

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetListVersionsConfigQuery)).
			WithArgs(configData.Name, params.Limit, params.Offset).
			WillReturnRows(rows)

		countColumn := []string{"count(*)"}
		totalRows := sqlmock.NewRows(countColumn)
		totalRows.AddRow(totalRowValue)

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetConfigVersionsTotalCountQuery)).
			WithArgs(configData.Name).
			WillReturnRows(totalRows)

		// When
		result, pagination, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.Nil(err)
		s.Equal(expectedResponse, result[0], "Result should be equal")
		s.LessOrEqual(int(params.Limit), len(result), "Result data should be less or equal as limit")
		s.Equal(expectedPagination, pagination, "Pagination should be equal")
	})
}

func (s *configsRepoSuite) TestConfigs_GetListVersion_ErrNotFound() {
	s.Run("Test Get Config List Version by config-name return Err Not Found", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetListConfigVersionsRequest{
			Name:   "test-config",
			Limit:  1,
			Offset: 0,
		}

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetListVersionsConfigQuery)).
			WithArgs(params.Name, params.Limit, params.Offset).
			WillReturnError(sql.ErrNoRows)

		// When
		result, pagination, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.Equal(sql.ErrNoRows, err, "Should return ErrNoRows")
		s.Nil(result)
		s.Nil(pagination)
	})
}

func (s *configsRepoSuite) TestConfigs_GetListVersion_ErrDB() {
	s.Run("Test Get Config List Version by config-name return Err DB", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetListConfigVersionsRequest{
			Name:   "test-config",
			Limit:  1,
			Offset: 0,
		}
		mockErr := testutil.ErrUnexpected

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetListVersionsConfigQuery)).
			WithArgs(params.Name, params.Limit, params.Offset).
			WillReturnError(mockErr)

		// When
		result, pagination, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
		s.Nil(result)
		s.Nil(pagination)
	})
}

func (s *configsRepoSuite) TestConfigs_GetListVersionTotalRows_ErrDB() {
	s.Run("Test Get Config List Total Rows by config-name return Err DB", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetListConfigVersionsRequest{
			Name:   "test-config",
			Limit:  1,
			Offset: 0,
		}
		mockErr := testutil.ErrUnexpected

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetListVersionsConfigQuery)).
			WithArgs(params.Name, params.Limit, params.Offset).
			WillReturnError(mockErr)

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.GetConfigVersionsTotalCountQuery)).
			WithArgs(params.Name).
			WillReturnError(mockErr)

		// When
		result, pagination, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
		s.Nil(result)
		s.Nil(pagination)
	})
}
