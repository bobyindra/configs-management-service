package auth_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	repoMock "github.com/bobyindra/configs-management-service/module/configuration/internal/repository/mock"
	sessionUsecase "github.com/bobyindra/configs-management-service/module/configuration/internal/usecase/auth"
	encryptMock "github.com/bobyindra/configs-management-service/module/configuration/util/mock"
)

type authUsecaseSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject usecase.SessionUsecase

	encryption *encryptMock.MockEncryption
	userRepo   *repoMock.MockUserRepository
}

func (s *authUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.encryption = encryptMock.NewMockEncryption(ctrl)
	s.userRepo = repoMock.NewMockUserRepository(ctrl)
	s.subject = sessionUsecase.NewSessionUscs(s.encryption, s.userRepo)
}

func (s *authUsecaseSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestSessionUsecaseSuite(t *testing.T) {
	suite.Run(t, new(authUsecaseSuite))
}
