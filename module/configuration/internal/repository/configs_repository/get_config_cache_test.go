package configs_repository_test

import (
	"context"
	"fmt"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/test"
	"github.com/redis/go-redis/v9"
)

func (s *configsRepoSuite) TestConfigs_GetConfigCache_Success() {
	s.Run("Get Config Cache - Success", func() {
		// Given
		ctx := context.TODO()
		cfgData := test.BuildConfigData()
		key := fmt.Sprintf("configs-%s", cfgData.Name)

		// marshal cfgData to json
		jsonData, err := util.ConvertAnyValueToJsonString(cfgData)

		// mock
		s.cacheMock.ExpectGet(key).SetVal(*jsonData)

		// When
		resp, err := s.subject.GetConfigCache(ctx, cfgData.Name)

		// Then
		s.Nil(err, "Error should be nil")
		s.Equal(cfgData, resp, "Response should be equal to cfgData")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
	})
}

func (s *configsRepoSuite) TestConfigs_GetConfigCache_Error() {
	s.Run("Get Config Cache - Unexpected Error", func() {
		// Given
		ctx := context.TODO()
		cfgName := "test-config"
		key := fmt.Sprintf("configs-%s", cfgName)
		mockErr := testutil.ErrUnexpected

		// mock
		s.cacheMock.ExpectGet(key).SetErr(mockErr)

		// When
		resp, err := s.subject.GetConfigCache(ctx, cfgName)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
		s.Nil(resp, "Response data should be nil")
	})

	s.Run("Get Config Cache - Redis Nil", func() {
		// Given
		ctx := context.TODO()
		cfgName := "test-config"
		key := fmt.Sprintf("configs-%s", cfgName)
		mockErr := redis.Nil

		// mock
		s.cacheMock.ExpectGet(key).SetErr(mockErr)

		// When
		resp, err := s.subject.GetConfigCache(ctx, cfgName)

		// Then
		s.Equal(mockErr, err, "Error should be equal to mockErr")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
		s.Nil(resp, "Response data should be nil")
	})

	s.Run("Get Config Cache - Unmarshal Error", func() {
		// Given
		ctx := context.TODO()
		cfgName := "test-config"
		key := fmt.Sprintf("configs-%s", cfgName)
		invalidData := make(chan int)

		// mock
		s.cacheMock.ExpectGet(key).SetVal(fmt.Sprint(invalidData))

		// When
		resp, err := s.subject.GetConfigCache(ctx, cfgName)

		// Then
		s.NotNil(err, "Error should be not nil")
		cacheMockErr := s.cacheMock.ExpectationsWereMet()
		s.Nil(cacheMockErr, "All Cache expectations should be met")
		s.Nil(resp, "Response data should be nil")
	})
}
