package configsusecase_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/golang/mock/gomock"
)

func (s *configsUsecaseSuite) TestConfigs_CreateConfig_Success() {
	s.Run("Create Config - Success", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
			ActorId:      1,
		}
		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().CreateConfig(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(nil)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
		s.Equal(config.Name, resp.Name, "Config name should be equal")
		s.Equal(config.ConfigValues, resp.ConfigValues, "Config name should be equal")
		s.Equal(config.ActorId, resp.ActorId, "Actor ID should be equal")
	})
}

func (s *configsUsecaseSuite) TestConfigs_CreateConfig_Err() {
	s.Run("Create Config - Schema Not Found Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
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

	s.Run("Create Config - Invalid Schema Type Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
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
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Create Config - Invalid Predefined Schema Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}
		schemaJSON := []byte(`{invalid schema`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.NotNil(err.Error(), "Error should be not nil")
	})

	s.Run("Create Config - Set DB Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}
		mockErr := testutil.ErrUnexpected
		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().CreateConfig(ctx, gomock.Any()).Return(mockErr)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(resp, "Response should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})

	s.Run("Create Config - Set Cache Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}
		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsDBRepo.EXPECT().CreateConfig(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(testutil.ErrUnexpected)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}
