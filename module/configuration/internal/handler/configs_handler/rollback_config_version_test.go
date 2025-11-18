package configshandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func (s *configsHandlerSuite) TestRollbackConfig_Success() {
	s.Run("Test Rollback Config - Success", func() {
		// Given
		params := &entity.Config{
			Name:    "wording-config",
			Version: 1,
		}

		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: "test config",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      1,
		}

		// mock
		s.configsUsecase.EXPECT().RollbackConfigVersionByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponse, nil)

		// When
		w := s.rollbackConfig(params, "rw")

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) TestRollbackConfig_Error() {
	s.Run("Test Rollback Config - Permission Denied", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}
		invalidRole := "no"

		// When
		w := s.rollbackConfig(params, invalidRole)

		// Then
		s.Equal(http.StatusForbidden, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrForbidden.Code, "Should contain error")
	})

	s.Run("Test Rollback Config - Body Invalid", func() {
		// Given
		params := "{invalid json"
		expectedErrorCode := "INTERNAL_ERROR"

		// When
		w := s.rollbackConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Rollback Config - Normalize Version Empty Error", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}

		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := s.rollbackConfig(params, "rw")

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Rollback Config - Normalize Name Empty Error", func() {
		// Given
		params := &entity.Config{
			Name:    "wording-config",
			Version: 1,
		}

		expectedErrorCode := "EMPTY_FIELD"

		// Set request data without giving :name value on the path
		gin.SetMode(gin.TestMode)
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(params)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/configs/:name/rollback", &buf)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		addClaim := &auth.AdditionalClaim{
			UserID: 1,
			Role:   "rw",
		}
		c.Set(middleware.ContextKeyAdditionalClaim, addClaim)

		// When
		s.subject.RollbackConfigVersion(c)

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Rollback Config - Error", func() {
		// Given
		params := &entity.Config{
			Name:    "wording-config",
			Version: 1,
		}

		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.configsUsecase.EXPECT().RollbackConfigVersionByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.rollbackConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})
}

func (s *configsHandlerSuite) rollbackConfig(body any, role string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/configs/:name/rollback", &buf)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "name", Value: "wording-config"},
	}
	addClaim := &auth.AdditionalClaim{
		UserID: 1,
		Role:   role,
	}
	c.Set(middleware.ContextKeyAdditionalClaim, addClaim)
	s.subject.RollbackConfigVersion(c)
	return w
}
