package configshandler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	usecaseMock "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/mock"
)

type configsHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject *configsHandler.ConfigsHandler

	configsUsecase *usecaseMock.MockConfigsManagementUsecase
}

func (s *configsHandlerSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.configsUsecase = usecaseMock.NewMockConfigsManagementUsecase(ctrl)
	s.subject = configsHandler.NewConfigsHandler(s.configsUsecase)
}

func (s *configsHandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsHandlerSuite(t *testing.T) {
	suite.Run(t, new(configsHandlerSuite))
}
