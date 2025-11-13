package configs_repository_test

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/mattn/go-sqlite3"
)

func (s *configsRepoSuite) TestConfigs_CreateConfig_Success() {
	s.Run("Test Create Config - Success", func() {
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

func (s *configsRepoSuite) TestConfigs_CreateConfig_Error() {
	s.Run("Test Create Config - Config Exist Err", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)

		sqliteErr := sqlite3.Error{ExtendedCode: sqlite3.ErrConstraintUnique}

		s.mock.ExpectQuery(regexp.QuoteMeta(configsRepo.CreateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.Version,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnError(sqliteErr)

		// When
		err = s.subject.CreateConfig(ctx, configData)

		// Then
		s.ErrorIs(err, entity.ErrConfigAlreadyExists, "Should return err config already exists")
	})

	s.Run("Test Create Config - Err DB", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)
		mockErr := testutil.ErrDB

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
