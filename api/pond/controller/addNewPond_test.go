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

	mocks "delos/mocks/api/pond"
)

func TestController_Pond_AddNewPond(t *testing.T) {
	var (
		endpoint = "/v1/ponds"
		payloads = payload.AddUpdatePond{
			Title:       "pond 2",
			FarmID:      2,
			Description: "pond 2 desc",
			Req:         "Add",
		}
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {
		requestBody, _ := json.Marshal(payloads)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("AddNewPond", payloads).
			Return(&model.Pond{
				ID: 1,
			}, nil)
		expectedStatusCode := http.StatusCreated

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.AddNewPond(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
