package util_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/bobyindra/configs-management-service/module/configuration/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_BuildSuccessResponse(t *testing.T) {
	t.Parallel()
	t.Run("Build Success Response return api response correct data", func(t *testing.T) {
		t.Parallel()

		// Given
		w := httptest.NewRecorder()

		dataRes := entity.ConfigResponse{
			ConfigValues: "Test",
		}
		metaRes := entity.PaginationResponse{}
		apiRes := util.APIResponse{
			Status: http.StatusOK,
			Data:   dataRes,
			Meta:   metaRes,
		}

		// When
		util.BuildSuccessResponse(w, apiRes)

		// Then
		assert.Equal(t, apiRes.Status, w.Code, "Response httpStatus should be equal")
		assert.Contains(t, w.Body.String(), dataRes.ConfigValues, "Should contains correct data")
	})
}

func TestUtil_BuildFailedResponse(t *testing.T) {
	t.Parallel()
	t.Run("Build Failed Response return api response with error detail provided", func(t *testing.T) {
		t.Parallel()

		// Given
		w := httptest.NewRecorder()
		errorRes := entity.ErrConfigAlreadyExists

		// When
		util.BuildFailedResponse(w, errorRes)

		// Then
		assert.Equal(t, errorRes.HttpCode, w.Code, "Response httpStatus should be equal")
		assert.Contains(t, w.Body.String(), errorRes.Code, "Should contains correct error")
	})

	t.Run("Build Failed Response return api response without provide error detail", func(t *testing.T) {
		t.Parallel()

		// Given
		w := httptest.NewRecorder()
		errorRes := testutil.ErrUnexpected

		// When
		util.BuildFailedResponse(w, errorRes)

		// Then
		assert.Equal(t, http.StatusInternalServerError, w.Code, "Response httpStatus should be equal")
		assert.Contains(t, w.Body.String(), errorRes.Error(), "Should contains correct error")
	})
}
