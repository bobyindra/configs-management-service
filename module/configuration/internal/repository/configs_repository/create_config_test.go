package configs_repository_test

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/internal/util"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/mattn/go-sqlite3"
)

func (s *configsRepoSuite) TestConfigs_CreateConfig_Success() {
	s.Run("Test Create Config by config-name return success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()

		scanColumn := []string{"id"}
		rows := sqlmock.NewRows(scanColumn)
		rows.AddRow(configData.Id)
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.CreateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.Version,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnRows(rows)

		// When
		err = s.subject.CreateConfig(ctx, configData)

		// Then
		s.Nil(err)
		sqlMockErr := s.mock.ExpectationsWereMet()
		s.Nil(sqlMockErr, "All DB expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_CreateConfig_ErrConstraintUnique() {
	s.Run("Test Create Config by config-name return Err Constraint Unique", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.CreateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.Version,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnError(sqlite3.ErrNo(sqlite3.ErrConstraintUnique))

		// When
		err = s.subject.CreateConfig(ctx, configData)

		// Then
		s.Equal(sqlite3.ErrNo(sqlite3.ErrConstraintUnique), err, "Should return ErrConfigAlreadyExists")
	})
}

func (s *configsRepoSuite) TestConfigs_CreateConfig_ErrDB() {
	s.Run("Test Create Config by config-name return Err DB", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)
		mockErr := testutil.ErrUnexpected

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.CreateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.Version,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnError(mockErr)

		// When
		err = s.subject.CreateConfig(ctx, configData)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
