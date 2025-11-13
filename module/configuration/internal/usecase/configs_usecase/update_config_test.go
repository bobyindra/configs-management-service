package configsusecase_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/golang/mock/gomock"
)

func (s *configsUsecaseSuite) TestConfigs_UpdateConfig_Success() {
	s.Run("Update Config - Success", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configRepo.EXPECT().UpdateConfigByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}

func (s *configsUsecaseSuite) TestConfigs_UpdateConfig_Err() {
	s.Run("Update Config - Err GetConfigByConfigName", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		mockErr := testutil.ErrUnexpected

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(nil, mockErr)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Err When Comparing Config Value", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: make(chan int),
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		mockErr := testutil.ErrJsonUnsupportedType

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Err Config Value Equal", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "config",
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		expectedErr := entity.ErrNoChangesFound

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(expectedErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Err", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		mockErr := testutil.ErrUnexpected

		// mock
		s.configRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configRepo.EXPECT().UpdateConfigByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(mockErr)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
