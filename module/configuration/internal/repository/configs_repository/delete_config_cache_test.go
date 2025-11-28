package configs_repository_test

import (
	"context"
	"fmt"

	"github.com/bobyindra/configs-management-service/internal/testutil"
)

func (s *configsCacheRepoSuite) TestConfigs_DeleteConfigCache_Success() {
	s.Run("Delete Config Cache - Success", func() {
		// Given
		ctx := context.TODO()
		cfgName := "test-config"
		key := fmt.Sprintf("configs-%s", cfgName)

		// mock
		s.cacheMock.ExpectDel(key).SetVal(int64(1))

		// When
		err := s.subject.DeleteConfigCache(ctx, cfgName)

		// Then
		s.Nil(err, "Error should be nil")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}

func (s *configsCacheRepoSuite) TestConfigs_DeleteConfigCache_Error() {
	s.Run("Delete Config Cache - Error", func() {
		// Given
		ctx := context.TODO()
		cfgName := "test-config"
		key := fmt.Sprintf("configs-%s", cfgName)
		mockErr := testutil.ErrUnexpected

		// mock
		s.cacheMock.ExpectDel(key).SetErr(mockErr)

		// When
		err := s.subject.DeleteConfigCache(ctx, cfgName)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}
