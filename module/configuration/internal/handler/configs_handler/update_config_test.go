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

func (s *configsHandlerSuite) TestUpdateConfig_Success() {
	s.Run("Test Update Config - Success", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: params.ConfigValues,
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      1,
		}

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsUsecase.EXPECT().UpdateConfigByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponse, nil)

		// When
		w := s.updateConfig(params, "rw")

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) TestUpdateConfig_Error() {
	s.Run("Test Update Config - Permission Denied", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}
		invalidRole := "no"

		// When
		w := s.updateConfig(params, invalidRole)

		// Then
		s.Equal(http.StatusForbidden, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrForbidden.Code, "Should contain error")
	})

	s.Run("Test Update Config - Body Invalid", func() {
		// Given
		params := "{invalid json"
		expectedErrorCode := "INTERNAL_ERROR"

		// When
		w := s.updateConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Update Config - Normalize Error", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := s.updateConfig(params, "rw")

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Update Config - Invalid Object Type", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		schemaJSON := []byte(`{
			"type": "integer",
			"minimum": 0,
			"maximum": 10000
		}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)

		// When
		w := s.updateConfig(params, "rw")

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrInvalidSchema.Message, "Should contain error")
	})

	s.Run("Test Update Config - Error", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsUsecase.EXPECT().UpdateConfigByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.updateConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})
}

func (s *configsHandlerSuite) updateConfig(body any, role string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/configs/:name", &buf)
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
	s.subject.UpdateConfig(c)
	return w
}
