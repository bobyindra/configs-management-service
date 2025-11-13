package configshandler_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
)

func (s *configsHandlerSuite) TestGetConfig_Success() {
	s.Run("Test Get Last Config - Success", func() {
		// Given
		params := &entity.GetConfigRequest{
			Name: "wording-config",
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
			ConfigValues: "test config",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetConfigByConfigName(gomock.Any(), params).Return(configResponse, nil)

		// When
		w := s.GetConfig("")

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})

	s.Run("Test Get Spesific Config Version - Success", func() {
		// Given
		params := &entity.GetConfigRequest{
			Name:    "wording-config",
			Version: 1,
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
			ConfigValues: "test config ver 1",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetConfigByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponse, nil)

		// When
		w := s.GetConfig("1")

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) TestGetConfig_Error() {
	s.Run("Test Get Config - Claim Error", func() {
		// Given
		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.GetConfig("")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains correct config value")
	})

	s.Run("Test Get Config - Forbidden", func() {
		// Given
		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "no",
			},
		}

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)

		// When
		w := s.GetConfig("")

		// Then
		s.Equal(http.StatusForbidden, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrForbidden.Code, "Should contains correct config value")
	})

	s.Run("Test Get Config - Normalize Error", func() {
		// Given
		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)

		// When
		w := s.GetConfig("abc")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains correct config value")
	})

	s.Run("Test Get Config - Error", func() {
		// Given
		params := &entity.GetConfigRequest{
			Name: "wording-config",
		}

		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetConfigByConfigName(gomock.Any(), params).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.GetConfig("")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) GetConfig(version string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	url := "/api/v1/configs/:name"
	if version != "" {
		url = url + "?version=" + version
	}

	req, _ := http.NewRequest(http.MethodPost, url, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "name", Value: "wording-config"},
	}
	s.subject.GetConfig(c)
	return w
}
