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
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
)

func (s *configsHandlerSuite) TestCreateConfig_Success() {
	s.Run("Test Create Config - Success", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: params.ConfigValues,
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}

		schemaJSON := []byte(`{
			"type": "string",
			"minLength": 5,
			"maxLength": 100
			}`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)
		s.configsUsecase.EXPECT().CreateConfig(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponse, nil)

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusCreated, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) TestCreateConfig_Error() {
	s.Run("Test Create Config - Permission Denied", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}
		invalidRole := "no"

		// When
		w := s.createConfig(params, invalidRole)

		// Then
		s.Equal(http.StatusForbidden, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrForbidden.Code, "Should contain error")
	})

	s.Run("Test Create Config - Body Invalid", func() {
		// Given
		params := "{invalid json"
		expectedErrorCode := "INTERNAL_ERROR"

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Create Config - Normalize Request config-value empty Error", func() {
		// Given
		params := &entity.Config{
			Name: "wording-config",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Create Config - Normalize Request name empty Error", func() {
		// Given
		params := &entity.Config{
			ConfigValues: "test values",
		}
		expectedErrorCode := "EMPTY_FIELD"

		// Set request data without giving :name value on the path
		gin.SetMode(gin.TestMode)
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(params)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/configs/:name", &buf)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		addClaim := &auth.AdditionalClaim{
			UserID: 1,
			Role:   "rw",
		}
		c.Set(middleware.ContextKeyAdditionalClaim, addClaim)

		// When
		s.subject.CreateConfigs(c)

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Create Config - Invalid Predefined Schema", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		expectedErrorCode := "INTERNAL_ERROR"
		schemaJSON := []byte(`{invalid schema`)

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(schemaJSON, nil)

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Create Config - Invalid object type", func() {
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
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusBadRequest, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrInvalidSchema.Message, "Should contain error")
	})

	s.Run("Test Create Config - Schema Not Found", func() {
		// Given
		params := &entity.Config{
			Name:         "wording-config",
			ConfigValues: "test values",
		}

		// mock
		s.schemaRegistry.EXPECT().GetSchemaByConfigName(gomock.Any()).Return(nil, entity.ErrConfigSchemaNotFound)

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusNotFound, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrConfigSchemaNotFound.Message, "Should contain error")
	})

	s.Run("Test Create Config - Error", func() {
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
		s.configsUsecase.EXPECT().CreateConfig(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.createConfig(params, "rw")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})
}

func (s *configsHandlerSuite) createConfig(body any, role string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/configs/:name", &buf)
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
	s.subject.CreateConfigs(c)
	return w
}
