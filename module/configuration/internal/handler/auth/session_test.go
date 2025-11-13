package auth_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/internal/handler/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	authMock "github.com/bobyindra/configs-management-service/module/configuration/internal/auth/mock"
	sessionHandler "github.com/bobyindra/configs-management-service/module/configuration/internal/handler/auth"
	usecaseMock "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/mock"
)

type authHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject *auth.SessionHandler

	auth           *authMock.MockAuth
	sessionUsecase *usecaseMock.MockSessionUsecase
}

func (s *authHandlerSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.auth = authMock.NewMockAuth(ctrl)
	s.sessionUsecase = usecaseMock.NewMockSessionUsecase(ctrl)
	s.subject = sessionHandler.NewSession(s.auth, s.sessionUsecase)
}

func (s *authHandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestSessionHandlerSuite(t *testing.T) {
	suite.Run(t, new(authHandlerSuite))
}
