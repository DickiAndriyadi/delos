package controller

import (
	"bytes"
	"delos/api/model"
	"delos/api/model/payload"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	mocks "delos/mocks/api/farm"
)

func TestController_Farm_AddNewFarm(t *testing.T) {
	var (
		endpoint = "/v1/farms"
		payloads = payload.AddUpdateFarm{
			Title:       "farm 2",
			Description: "farm 2 desc",
		}
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {
		requestBody, _ := json.Marshal(payloads)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("AddNewFarm", payloads).
			Return(&model.Farm{
				ID: 1,
			}, nil)
		expectedStatusCode := http.StatusCreated

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.AddNewFarm(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
