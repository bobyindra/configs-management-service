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

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().UpdateConfigByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().DeleteConfigCache(ctx, config.Name).Return(nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
		s.Equal(config.Name, resp.Name, "Config name should be equal")
		s.Equal(config.ConfigValues, resp.ConfigValues, "Config name should be equal")
		s.Equal(config.ActorId, resp.ActorId, "Actor ID should be equal")
	})
}

func (s *configsUsecaseSuite) TestConfigs_UpdateConfig_Err() {
	s.Run("Create Config - Schema Not Found Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}
		mockErr := entity.ErrConfigSchemaNotFound

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(nil, mockErr)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Invalid Schema Type Err", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}

		schemaJSON := []byte(`{
			"type": "integer",
			"minimum": 0,
			"maximum": 10000
		}`)
		mockErr := entity.ErrInvalidSchema

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Invalid Predefined Schema Err", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "value",
			ActorId:      1,
		}

		schemaJSON := []byte(`{invalid schema`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.NotNil(err.Error(), "Error should be not nil")
	})

	s.Run("Update Config - GetConfigByConfigName Err", func() {
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

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(nil, mockErr)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Comparing Config Value Err", func() {
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "test",
			ConfigValues: "config",
			ActorId:      1,
		}
		getParam := &entity.GetConfigRequest{
			Name: config.Name,
		}
		cfgResponse := &entity.ConfigResponse{
			Id:           5,
			Name:         getParam.Name,
			ConfigValues: make(chan int),
			ActorId:      1,
		}
		mockErr := testutil.ErrJsonUnsupportedType

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Config Value Equal Err", func() {
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

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(expectedErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Set DB Err", func() {
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

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().UpdateConfigByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(mockErr)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Update Config - Set Cache Err", func() {
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

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().GetConfigByConfigName(ctx, getParam).Return(cfgResponse, nil)
		s.configsDBRepo.EXPECT().UpdateConfigByConfigName(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().DeleteConfigCache(ctx, config.Name).Return(testutil.ErrUnexpected)

		// When
		resp, err := s.subject.UpdateConfigByConfigName(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}
