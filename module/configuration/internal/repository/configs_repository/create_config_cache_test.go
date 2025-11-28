package configs_repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
)

func (s *configsCacheRepoSuite) TestConfigs_CreateConfigCache_Success() {
	s.Run("Create Config Cache - Success", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		jsonData, err := json.Marshal(configData)
		key := fmt.Sprintf("configs-%s", configData.Name)
		ttl := 12 * time.Hour

		// mock
		s.cacheMock.ExpectSet(key, jsonData, ttl).SetVal("OK")

		// When
		err = s.subject.CreateConfigCache(ctx, configData)

		// Then
		s.Nil(err, "Error should be nil")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}

func (s *configsCacheRepoSuite) TestConfigs_CreateConfigCache_Error() {
	s.Run("Create Config Cache - Error", func() {
		// Given
		ctx := context.TODO()
		configData := test.BuildConfigData()
		jsonData, err := json.Marshal(configData)
		key := fmt.Sprintf("configs-%s", configData.Name)
		ttl := 12 * time.Hour
		mockErr := testutil.ErrUnexpected

		// mock
		s.cacheMock.ExpectSet(key, jsonData, ttl).SetErr(mockErr)

		// When
		err = s.subject.CreateConfigCache(ctx, configData)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}
