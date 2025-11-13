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

func (h *authHandlerSuite) TestAuthLogin_Success() {
	h.Run("Test Login - Success", func() {
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
		h.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(userResponse, nil)
		h.auth.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return(mockToken, nil)

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusOK, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), mockToken, "Should contains token")
	})
}

func (h *authHandlerSuite) TestAuthLogin_Error() {
	h.Run("Test Login - decode body Err", func() {
		// Given
		params := "{invalid json"
		expectedErrorCode := "INTERNAL_ERROR"

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), expectedErrorCode, "Should contains expected error message")
	})

	h.Run("Test Login - normalize request password Err", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	h.Run("Test Login - normalize request username Err", func() {
		// Given
		params := &entity.LoginRequest{
			Password: "testpassword",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	h.Run("Test Login - login Err", func() {
		// Given
		params := &entity.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}
		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		h.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(nil, testutil.ErrUnexpected)

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	h.Run("Test Login - generate token Err", func() {
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
		h.sessionUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(userResponse, nil)
		h.auth.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", testutil.ErrUnexpected)

		// When
		w := h.doLogin(params)

		// Then
		h.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		h.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})
}

func (h *authHandlerSuite) doLogin(body any) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(http.MethodPost, "/login", &buf)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	h.subject.Login(c)
	return w
}
