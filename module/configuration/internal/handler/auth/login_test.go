package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func (s *authHandlerSuite) TestAuthLogin_Success() {
	s.Run("Test Login - Success", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}

		userResponse := &entity.LoginResponse{
			UserID: 1,
			Role:   "rw",
		}

		mockToken := "mocktoken"

		// mock
		s.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(userResponse, nil)
		s.auth.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return(mockToken, nil)

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), mockToken, "Should contains token")
	})
}

func (s *authHandlerSuite) TestAuthLogin_Error() {
	s.Run("Test Login - decode body Err", func() {
		// Given
		params := "{invalid json"
		expectedErrorCode := "INTERNAL_ERROR"

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains expected error message")
	})

	s.Run("Test Login - normalize request password Err", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Login - normalize request username Err", func() {
		// Given
		params := &entity.LoginRequest{
			Password: "testpassword",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Login - login Err", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}
		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Login - generate token Err", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}
		expectedErrorCode := "INTERNAL_ERROR"
		userResponse := &entity.LoginResponse{
			UserID: 1,
			Role:   "rw",
		}

		// mock
		s.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(userResponse, nil)
		s.auth.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", testutil.ErrUnexpected)

		// When
		w := s.doLogin(params)

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})
}

func (s *authHandlerSuite) doLogin(body any) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", &buf)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	s.subject.Login(c)
	return w
}
