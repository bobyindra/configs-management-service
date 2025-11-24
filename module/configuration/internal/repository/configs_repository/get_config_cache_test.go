package configs_repository_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (s *configsRepoSuite) TestConfigs_GetConfigCache_Success() {
	s.Run("Get Config Cache - Success", func() {
		// Given
		ctx := context.TODO()
		key := "test-config"
		cfgData := entity.Config{}

		// mock
		s.cacheMock.ExpectGetSet(key, &cfgData).SetVal("Value")

		// When
		resp, err := s.subject.GetConfigCache(ctx, key)

		// Then
		s.Nil(err, "Error should be nil")
		s.Equal(&cfgData, resp, "Response should be equal to cfgData")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_GetConfigCache_Error() {
	s.Run("Get Config Cache - Error", func() {
		// Given
		ctx := context.TODO()
		key := "test-config"
		mockErr := testutil.ErrUnexpected

		// mock
		s.cacheMock.ExpectGetSet(key, &entity.Config{}).SetErr(mockErr)

		// When
		resp, err := s.subject.GetConfigCache(ctx, key)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
		s.Nil(resp, "Response data should be nil")
	})
}
