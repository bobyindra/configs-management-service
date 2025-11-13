package configsusecase_test

import (
	"context"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

func (s *configsUsecaseSuite) TestConfigs_GetConfigVersionList_Success() {
	s.Run("Get Config Version List - Success", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetListConfigVersionsRequest{
			Name:   "test",
			Limit:  2,
			Offset: 0,
		}

		configs := []*entity.ConfigResponse{}
		pagination := &entity.PaginationResponse{}

		// mock
		s.configRepo.EXPECT().GetListVersionsByConfigName(ctx, params).Return(configs, pagination, nil)

		// When
		resp, pg, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.Nil(err, "Error should be nil")
		s.NotNil(resp, "Response should in place")
		s.NotNil(pg, "Pagination should in place")
	})
}

func (s *configsUsecaseSuite) TestConfigs_GetConfigVersionList_Err() {
	s.Run("Get Config Version List - Err", func() {
		// Given
		ctx := context.TODO()
		params := &entity.GetListConfigVersionsRequest{
			Name:   "test",
			Limit:  2,
			Offset: 0,
		}
		mockErr := testutil.ErrUnexpected

		// mock
		s.configRepo.EXPECT().GetListVersionsByConfigName(ctx, params).Return(nil, nil, mockErr)

		// When
		resp, pg, err := s.subject.GetListVersionsByConfigName(ctx, params)

		// Then
		s.Nil(resp, "Response should be nil")
		s.Nil(pg, "Pagination should be nil")
		s.EqualError(mockErr, err.Error(), "Error should be equal")
	})
}
