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

func (s *configsHandlerSuite) TestGetListConfigVersion_Success() {
	s.Run("Test Get List Config Version - Success", func() {
		params := &entity.GetListConfigVersionsRequest{
			Name:   "wording-config",
			Limit:  2,
			Offset: 0,
		}

		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		configResponses := []*entity.ConfigResponse{}
		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: "test config",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}
		configResponses = append(configResponses, configResponse)

		pagination := &entity.PaginationResponse{}

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetListVersionsByConfigName(gomock.Any(), params).Return(configResponses, pagination, nil)

		// When
		w := s.GetListConfigVersions("")

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})

	s.Run("Test Get List Config Version with false limit - Success", func() {
		params := &entity.GetListConfigVersionsRequest{
			Name: "wording-config",
		}

		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		configResponses := []*entity.ConfigResponse{}
		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: "test config",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}
		configResponses = append(configResponses, configResponse)

		pagination := &entity.PaginationResponse{}

		url := "/api/v1/configs/:name/versions?limit=abcd&offset=0"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetListVersionsByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponses, pagination, nil)

		// When
		w := s.GetListConfigVersions(url)

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})

	s.Run("Test Get List Config Version with false offset - Success", func() {
		params := &entity.GetListConfigVersionsRequest{
			Name: "wording-config",
		}

		jwtResponse := &auth.ConfigsJWTClaim{
			RegisteredClaims: jwt.RegisteredClaims{},
			AdditionalClaim: auth.AdditionalClaim{
				UserID: 1,
				Role:   "rw",
			},
		}

		configResponses := []*entity.ConfigResponse{}
		configResponse := &entity.ConfigResponse{
			Id:           1,
			Name:         params.Name,
			ConfigValues: "test config",
			Version:      1,
			CreatedAt:    time.Now().UTC(),
			ActorId:      jwtResponse.UserID,
		}
		configResponses = append(configResponses, configResponse)

		pagination := &entity.PaginationResponse{}

		url := "/api/v1/configs/:name/versions?limit=2&offset=abcd"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(jwtResponse, nil)
		s.configsUsecase.EXPECT().GetListVersionsByConfigName(gomock.Any(), gomock.AssignableToTypeOf(params)).Return(configResponses, pagination, nil)

		// When
		w := s.GetListConfigVersions(url)

		// Then
		s.Equal(http.StatusOK, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), configResponse.ConfigValues, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) TestGetListConfigVersion_Error() {
	s.Run("Test Get List Config Version - Claim Error", func() {
		// Given
		expectedErrorCode := "INTERNAL_ERROR"

		// mock
		s.auth.EXPECT().ValidateClaim(gomock.Any(), gomock.Any()).Return(nil, testutil.ErrUnexpected)

		// When
		w := s.GetListConfigVersions("")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contain error")
	})

	s.Run("Test Get List Config Version - Permission Denied", func() {
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
		w := s.GetListConfigVersions("")

		// Then
		s.Equal(http.StatusForbidden, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), entity.ErrForbidden.Code, "Should contains correct config value")
	})

	s.Run("Test Get List Config Version - Error", func() {
		// Given
		params := &entity.GetListConfigVersionsRequest{
			Name:   "wording-config",
			Limit:  2,
			Offset: 0,
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
		s.configsUsecase.EXPECT().GetListVersionsByConfigName(gomock.Any(), params).Return(nil, nil, testutil.ErrUnexpected)

		// When
		w := s.GetListConfigVersions("")

		// Then
		s.Equal(http.StatusInternalServerError, w.Code, "Status code should be equal")
		s.Contains(w.Body.String(), expectedErrorCode, "Should contains correct config value")
	})
}

func (s *configsHandlerSuite) GetListConfigVersions(customURL string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	url := "/api/v1/configs/:name/versions?limit=2&offset=0"
	if customURL != "" {
		url = customURL
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "name", Value: "wording-config"},
	}
	s.subject.GetConfigVersions(c)
	return w
}
