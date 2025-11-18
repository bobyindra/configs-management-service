package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	authMock "github.com/bobyindra/configs-management-service/module/configuration/internal/auth/mock"
)

type middlewareSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	subject *middleware.Middleware
	auth    *authMock.MockAuth
}

func (s *middlewareSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.ctrl = ctrl
	s.auth = authMock.NewMockAuth(ctrl)
	s.subject = middleware.NewMiddleware(s.auth)
}

func (s *middlewareSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(middlewareSuite))
}

func (s *middlewareSuite) TestMiddleware_ValidateSession() {
	s.Run("Validate Session - Success", func() {
		// Given
		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)

		// When
		c, _ := s.validateSession()
		ctxClaim, exist := c.Get(middleware.ContextKeyAdditionalClaim)
		claim, ok := ctxClaim.(*auth.AdditionalClaim)

		// Then
		s.True(exist, "Context Claim should be present")
		s.True(ok, "Claim format should be ok")
		s.Equal(jwtResponse.AdditionalClaim.UserID, claim.UserID, "User ID should be equal")
		s.Equal(jwtResponse.AdditionalClaim.Role, claim.Role, "Role should be equal")
	})

	s.Run("Validate Session - Error", func() {
		// Given
		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(nil, testutil.ErrUnexpected)

		// When
		_, w := s.validateSession()

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains correct config value")
	})
}

func (s *middlewareSuite) validateSession() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	s.subject.ValidateSession(c)
	return c, w
}
