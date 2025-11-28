package configsusecase_test

import (
	"context"
	"fmt"
	"time"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/golang/mock/gomock"
)

func (s *configsUsecaseSuite) TestConfigs_RollbackConfig_Success() {
	s.Run("Rollback Config - Success", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().RollbackConfigVersionByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(nil)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}

func (s *configsUsecaseSuite) TestConfigs_RollbackConfig_Err() {
	s.Run("Rollback Config - Err GetConfigByConfigName", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		mockErr := testutil.ErrUnexpected

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(nil, mockErr)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Rollback Config - Err Version is higher than lastest version", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 3,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		expectedErr := entity.ErrConfigVersionNotFound(fmt.Sprint(config.Version))

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(expectedErr, err.Error(), "Error should be equal")
	})

	s.Run("Rollback Config - Err Version is the same with lastest version", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 2,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		expectedErr := entity.ErrRollbackNotAllowed

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(expectedErr, err.Error(), "Error should be equal")
	})

	s.Run("Rollback Config - Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)
		mockErr := testutil.ErrUnexpected

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().RollbackConfigVersionByConfigName(ctx, gomock.Any()).Return(mockErr)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Rollback Config - Set Cache Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:    "test",
			Version: 1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := BuildConfigResponse(getParam)

		// mock
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().RollbackConfigVersionByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(testutil.ErrUnexpected)

		// When
		resp, err := s.subject.RollbackConfigVersionByConfigName(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}

func BuildConfigResponse(params *entity.GetConfigRequest) *entity.ConfigResponse {
	return &entity.ConfigResponse{
		Id:           5,
		Name:         params.Name,
		ConfigValues: "config",
		Version:      2,
		CreatedAt:    time.Now().UTC(),
		ActorId:      1,
	}
}
