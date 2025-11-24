package configs_repository_test

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bobyindra/configs-management-service/internal/testutil"
	configsRepo "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/configs_repository"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
)

func (s *configsRepoSuite) TestConfigs_RollbackConfig_Success() {
	s.Run("Test Rollback Config version return success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()

		scanColumn := []string{"id", "config_values", "version"}
		rows := sqlmock.NewRows(scanColumn)
		rows.AddRow(configData.Id, configData.ConfigValues, configData.Version)

		s.sqlMock.ExpectQuery(regexp.QuoteMeta(configsRepo.RollbackConfigVersionQuery)).
			WithArgs(
				configData.Name,
				configData.CreatedAt,
				configData.ActorId,
				configData.Version,
			).
			WillReturnRows(rows)

		// When
		err := s.subject.RollbackConfigVersionByConfigName(ctx, configData)

		// Then
		s.Nil(err)
		sqlMockErr := s.sqlMock.ExpectationsWereMet()
		s.Nil(sqlMockErr, "All DB expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_RollbackConfig_ErrDB() {
	s.Run("Test Rollback Config version return err db", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		mockErr := testutil.ErrUnexpected

		s.sqlMock.ExpectQuery(regexp.QuoteMeta(configsRepo.RollbackConfigVersionQuery)).
			WithArgs(
				configData.Name,
				configData.CreatedAt,
				configData.ActorId,
				configData.Version,
			).
			WillReturnError(mockErr)

		// When
		err := s.subject.RollbackConfigVersionByConfigName(ctx, configData)

		// Then
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
