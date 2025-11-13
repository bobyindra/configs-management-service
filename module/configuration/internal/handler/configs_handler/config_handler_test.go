package configshandler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	authMock "github.com/bobyindra/configs-management-service/module/configuration/internal/auth/mock"
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	usecaseMock "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/mock"
)

type configsHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject *configsHandler.ConfigsHandler

	auth           *authMock.MockAuth
	configsUsecase *usecaseMock.MockConfigsManagementUsecase
}

func (h *configsHandlerSuite) SetupTest() {
	ctrl := gomock.NewController(h.T())

	h.ctrl = ctrl
	h.auth = authMock.NewMockAuth(ctrl)
	h.configsUsecase = usecaseMock.NewMockConfigsManagementUsecase(ctrl)
	h.subject = configsHandler.NewConfigsHandler(h.auth, h.configsUsecase)
}

func (h *configsHandlerSuite) TearDownTest() {
	h.ctrl.Finish()
}

func TestConfigsHandlerSuite(t *testing.T) {
	suite.Run(t, new(configsHandlerSuite))
}
