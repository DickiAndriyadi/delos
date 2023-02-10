package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"delos/api/model"
	"delos/api/model/payload"
	mocks "delos/mocks/api/farm"
)

func TestController_Farm_UpdateFarm(t *testing.T) {
	var (
		id       string
		endpoint = fmt.Sprintf("/v1/farms/%s", id)
		payloads = payload.AddUpdateFarm{
			Title:       "farm 3",
			Description: "farm 3 desc",
		}
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {
		requestBody, _ := json.Marshal(payloads)

		e := echo.New()
		req := httptest.NewRequest(echo.PUT, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("UpdateFarm", payloads, id).
			Return(&model.Farm{
				ID: 1,
			}, nil)
		expectedStatusCode := http.StatusOK

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.UpdateFarm(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
