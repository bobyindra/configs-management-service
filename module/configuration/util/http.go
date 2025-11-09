package util

import (
	"encoding/json"
	"net/http"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
)

type APIResponse struct {
	Status  int                `json:"-"`
	Message string             `json:"message,omitempty"`
	Meta    any                `json:"meta,omitempty"`
	Data    any                `json:"omitempty"`
	Error   entity.ErrorDetail `json:"error,omitempty"`
}

func BuildSuccessResponse(w http.ResponseWriter, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	apiResponse := APIResponse{}
	if resp.Meta != nil {
		apiResponse.Meta = resp.Meta
	}
	if resp.Data != nil {
		apiResponse.Data = resp.Data
	}

	json.NewEncoder(w).Encode(apiResponse)
}

func BuildFailedResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	b, ok := err.(*entity.ErrorDetail)
	if ok {
		w.WriteHeader(b.HttpCode)

		json.NewEncoder(w).Encode(APIResponse{
			Error: *b,
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(APIResponse{
		Error: *entity.WrapError(err),
	})
}
