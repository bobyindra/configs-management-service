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

	configRepo *repoMock.MockConfigsManagementRepository
}

func (s *configsUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.configRepo = repoMock.NewMockConfigsManagementRepository(ctrl)
	s.subject = configUsecase.NewConfigsUsecase(s.configRepo)
}

func (s *configsUsecaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsUsecaseSuite(t *testing.T) {
	suite.Run(t, new(configsUsecaseSuite))
}
