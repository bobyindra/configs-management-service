package configs_repository_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
)

func (s *configsRepoSuite) TestConfigs_DeleteConfigCache_Success() {
	s.Run("Delete Config Cache - Success", func() {
		// Given
		ctx := context.TODO()
		key := "test-config"

		// mock
		s.cacheMock.ExpectDel(key).SetVal(int64(1))

		// When
		err := s.subject.DeleteConfigCache(ctx, key)

		// Then
		s.Nil(err, "Error should be nil")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_DeleteConfigCache_Error() {
	s.Run("Delete Config Cache - Error", func() {
		// Given
		ctx := context.TODO()
		key := "test-config"
		mockErr := testutil.ErrUnexpected

		// mock
		s.cacheMock.ExpectDel(key).SetErr(mockErr)

		// When
		err := s.subject.DeleteConfigCache(ctx, key)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}
