package configshandler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	authMock "github.com/bobyindra/configs-management-service/module/configuration/internal/auth/mock"
	configsHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/configs_handler"
	usecaseMock "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/mock"
	schemaMock "github.com/bobyindra/configs-management-service/module/configuration/schema/mock"
)

type configsHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject *configsHandler.ConfigsHandler

	auth           *authMock.MockAuth
	configsUsecase *usecaseMock.MockConfigsManagementUsecase
	schemaRegistry *schemaMock.MockSchemaRegistry
}

func (s *configsHandlerSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.auth = authMock.NewMockAuth(ctrl)
	s.configsUsecase = usecaseMock.NewMockConfigsManagementUsecase(ctrl)
	s.schemaRegistry = schemaMock.NewMockSchemaRegistry(ctrl)
	s.subject = configsHandler.NewConfigsHandler(s.auth, s.configsUsecase, s.schemaRegistry)
}

func (s *configsHandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestConfigsHandlerSuite(t *testing.T) {
	suite.Run(t, new(configsHandlerSuite))
}
