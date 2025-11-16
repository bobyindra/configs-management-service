package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegration_CreateConfig(t *testing.T) {
	t.Run("Create first config return success", func(t *testing.T) {
		// Given
		body := `{
			"config_values": true
		}`
		req, _ := http.NewRequest(
			http.MethodPost,
			server.URL+"/api/v1/configs/bca-enabled",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusCreated, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Equal(t, true, apiResponse.Data.ConfigValues)
		assert.Equal(t, uint16(1), apiResponse.Data.Version)
	})

	t.Run("Create the same config return config exists", func(t *testing.T) {
		// Given
		body := `{
			"config_values": true
		}`
		req, _ := http.NewRequest(
			http.MethodPost,
			server.URL+"/api/v1/configs/bca-enabled",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusConflict, resp.StatusCode)
		require.Nil(t, apiResponse.Data)
		require.NotNil(t, apiResponse.Error)
		assert.Equal(t, entity.ErrConfigAlreadyExists.Code, apiResponse.Error.Code)
	})
}

func TestIntegration_UpdateConfig(t *testing.T) {
	t.Run("Update with the same value return changes not found", func(t *testing.T) {
		// Given
		body := `{
			"config_values": true
		}`
		req, _ := http.NewRequest(
			http.MethodPut,
			server.URL+"/api/v1/configs/bca-enabled",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		require.Nil(t, apiResponse.Data)
		require.NotNil(t, apiResponse.Error)
		assert.Equal(t, entity.ErrNoChangesFound.Code, apiResponse.Error.Code)
	})

	t.Run("Update with the different value return success", func(t *testing.T) {
		// Given
		body := `{
			"config_values": false
		}`
		req, _ := http.NewRequest(
			http.MethodPut,
			server.URL+"/api/v1/configs/bca-enabled",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Equal(t, false, apiResponse.Data.ConfigValues)
		assert.Equal(t, uint16(2), apiResponse.Data.Version)
	})
}

func TestIntegration_RollbackConfig(t *testing.T) {
	t.Run("Rollback to the same latest version return not allowed", func(t *testing.T) {
		// Given
		body := `{
			"version": 2
		}`
		req, _ := http.NewRequest(
			http.MethodPost,
			server.URL+"/api/v1/configs/bca-enabled/rollback",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		require.Nil(t, apiResponse.Data)
		require.NotNil(t, apiResponse.Error)
		assert.Equal(t, entity.ErrRollbackNotAllowed.Code, apiResponse.Error.Code)
	})

	t.Run("Rollback to lower version return success", func(t *testing.T) {
		// Given
		body := `{
			"version": 1
		}`
		req, _ := http.NewRequest(
			http.MethodPost,
			server.URL+"/api/v1/configs/bca-enabled/rollback",
			bytes.NewBufferString(body),
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Equal(t, true, apiResponse.Data.ConfigValues)
		assert.Equal(t, uint16(3), apiResponse.Data.Version)
	})
}

func TestIntegration_GetConfig(t *testing.T) {
	t.Run("Get Latest Config", func(t *testing.T) {
		// Given
		req, _ := http.NewRequest(
			http.MethodGet,
			server.URL+"/api/v1/configs/bca-enabled",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Equal(t, uint16(3), apiResponse.Data.Version)
	})

	t.Run("Get Config Version 1", func(t *testing.T) {
		// Given
		req, _ := http.NewRequest(
			http.MethodGet,
			server.URL+"/api/v1/configs/bca-enabled?version=1",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  *entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail    `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Equal(t, uint16(1), apiResponse.Data.Version)
	})

}

func TestIntegration_GetListConfigVersion(t *testing.T) {
	t.Run("Get All List version with default limit 10", func(t *testing.T) {
		// Given
		req, _ := http.NewRequest(
			http.MethodGet,
			server.URL+"/api/v1/configs/bca-enabled/versions",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  []*entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail      `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Len(t, apiResponse.Data, 3)
	})

	t.Run("Get All List version with limit 2", func(t *testing.T) {
		// Given
		req, _ := http.NewRequest(
			http.MethodGet,
			server.URL+"/api/v1/configs/bca-enabled/versions?limit=2&offset=0",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")

		// When
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		var apiResponse struct {
			Data  []*entity.ConfigResponse `json:"data"`
			Error *entity.ErrorDetail      `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&apiResponse)

		// Then
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Nil(t, apiResponse.Error)
		require.NotNil(t, apiResponse.Data)
		assert.Len(t, apiResponse.Data, 2)
	})
}
