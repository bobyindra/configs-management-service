package configsusecase_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (s *configsUsecaseSuite) TestConfigs_GetConfigByConfigName_Success() {
	s.Run("Get Config by Config Name - Success", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetConfigRequest{
			Name:    "test",
			Version: 1,
		}

		config := &entity.ConfigResponse{}

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, params).Return(config, nil)

		// When
		resp, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}

func (s *configsUsecaseSuite) TestConfigs_GetConfigByConfigName_Err() {
	s.Run("Get Config by Config Name - Err", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetConfigRequest{
			Name:    "test",
			Version: 1,
		}
		mockErr := testutil.ErrUnexpected

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, params).Return(nil, mockErr)

		// When
		resp, err := s.subject.GetConfigByConfigName(ctx, params)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
