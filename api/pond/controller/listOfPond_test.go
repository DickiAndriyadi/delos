package controller

import (
	"delos/api/model"
	"delos/api/pond/filter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	mocks "delos/mocks/api/pond"
)

func TestController_Pond_ListOfPond(t *testing.T) {
	var (
		endpoint = "/v1/ponds"
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		filter := filter.FilterListOfPond{
			FarmID: "",
		}

		serviceMock := new(mocks.Service)
		serviceMock.On("ListOfPond", &filter).
			Return(&model.Ponds{
				{
					ID: 1,
				},
			}, nil)
		expectedStatusCode := http.StatusOK

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.ListOfPond(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
