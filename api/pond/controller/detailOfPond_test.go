package controller

import (
	"delos/api/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	mocks "delos/mocks/api/pond"
)

func TestController_Pond_DetailOfPond(t *testing.T) {
	var (
		id       string
		endpoint = fmt.Sprintf("/v1/ponds/%s", id)
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("DetailOfPond", id).
			Return(&model.Pond{ID: 1}, nil)
		expectedStatusCode := http.StatusOK

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.DetailOfPond(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
