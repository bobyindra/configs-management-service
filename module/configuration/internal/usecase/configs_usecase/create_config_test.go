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
		config := &entity.Config{}

		// mock
		s.configsDBRepo.EXPECT().CreateConfig(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(nil)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}

func (s *configsUsecaseSuite) TestConfigs_CreateConfig_Err() {
	s.Run("Create Config - Err", func() {
		// Given
		ctx := context.TODO()
		config := &entity.Config{}
		mockErr := testutil.ErrUnexpected

		// mock
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
		config := &entity.Config{}

		// mock
		s.configsDBRepo.EXPECT().CreateConfig(ctx, gomock.AssignableToTypeOf(config)).Return(nil)
		s.configsCacheRepo.EXPECT().CreateConfigCache(ctx, gomock.AssignableToTypeOf(config)).Return(testutil.ErrUnexpected)

		// When
		resp, err := s.subject.CreateConfig(ctx, config)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
	})
}
