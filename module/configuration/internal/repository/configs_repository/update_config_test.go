package configs_repository_test

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/internal/util"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
)

func (s *configsRepoSuite) TestConfigs_UpdateConfig_Success() {
	s.Run("Test Update Config version return success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()

		scanColumn := []string{"id", "version"}
		rows := sqlmock.NewRows(scanColumn)
		rows.AddRow(configData.Id, configData.Version)
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)

		s.sqlMock.ExpectQuery(regexp.QuoteMeta(configsRepo.UpdateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnRows(rows)

		// When
		err = s.subject.UpdateConfigByConfigName(ctx, configData)

		// Then
		s.Nil(err)
		sqlMockErr := s.sqlMock.ExpectationsWereMet()
		s.Nil(sqlMockErr, "All DB expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_UpdateConfig_ErrDB() {
	s.Run("Test Update Config version return err db", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		rowCfgValue, err := util.ConvertAnyValueToJsonString(configData.ConfigValues)
		mockErr := testutil.ErrUnexpected

		s.sqlMock.ExpectQuery(regexp.QuoteMeta(configsRepo.UpdateConfigQuery)).
			WithArgs(
				configData.Name,
				rowCfgValue,
				configData.CreatedAt,
				configData.ActorId,
			).
			WillReturnError(mockErr)

		// When
		err = s.subject.UpdateConfigByConfigName(ctx, configData)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
