package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	mocks "delos/mocks/api/farm"
)

func TestController_Farm_DeleteFarm(t *testing.T) {
	var (
		id       string
		endpoint = fmt.Sprintf("/v1/farms/%s", id)
	)

	t.Run("TestCase #1 : Positive", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		serviceMock := new(mocks.Service)
		serviceMock.On("DeleteFarm", id).
			Return("success delete data!", nil)
		expectedStatusCode := http.StatusOK

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := NewController(serviceMock)
		_ = controller.DeleteFarm(c)
		assert.Equal(t, expectedStatusCode, rec.Code)
	})
}
