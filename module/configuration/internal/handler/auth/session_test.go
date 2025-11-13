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

func (h *authHandlerSuite) SetupTest() {
	ctrl := gomock.NewController(h.T())

	h.ctrl = ctrl
	h.auth = authMock.NewMockAuth(ctrl)
	h.sessionUsecase = usecaseMock.NewMockSessionUsecase(ctrl)
	h.subject = sessionHandler.NewSession(h.auth, h.sessionUsecase)
}

func (h *authHandlerSuite) TearDownTest() {
	h.ctrl.Finish()
}

func TestSessionHandlerSuite(t *testing.T) {
	suite.Run(t, new(authHandlerSuite))
}
