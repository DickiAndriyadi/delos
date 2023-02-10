package controller

import (
	"delos/api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	mocks "delos/mocks/api/farm"
)

func TestController_Farm_ListOfFarm(t *testing.T) {
	var (
		endpoint = "/v1/farms"
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("ListOfFarm").
			Return(&model.Farms{
				{
					ID: 1,
				},
			}, nil)
		expectedStatusCode := http.StatusOK

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.ListOfFarm(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
