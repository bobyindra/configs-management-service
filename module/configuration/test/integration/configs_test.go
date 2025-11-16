package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/helper"
)

func TestIntegration_CreateConfig(t *testing.T) {
	// Given
	body := `{
		"config_values": true
	}`
	req, _ := http.NewRequest(http.MethodPost, server.URL+"/api/v1/configs/bca-enabled", bytes.NewBufferString(body))
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// When
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	apiResponse := helper.APIResponse{}
	json.NewDecoder(resp.Body).Decode(&apiResponse)

	// Then
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201, got %d", resp.StatusCode)
	}
}
