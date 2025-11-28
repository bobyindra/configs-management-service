package configsusecase_test

import (
	"testing"

	repoMock "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/mock"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	configUsecase "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/configs_usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type configsUsecaseSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject usecase.ConfigsManagementUsecase

	configsDBRepo    *repoMock.MockConfigsManagementDBRepository
	configsCacheRepo *repoMock.MockConfigsManagementCacheRepository
}

func (s *configsUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.configsDBRepo = repoMock.NewMockConfigsManagementDBRepository(ctrl)
	s.configsCacheRepo = repoMock.NewMockConfigsManagementCacheRepository(ctrl)
	s.subject = configUsecase.NewConfigsUsecase(s.configsDBRepo, s.configsCacheRepo)
}

func (s *configsUsecaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsUsecaseSuite(t *testing.T) {
	suite.Run(t, new(configsUsecaseSuite))
}
