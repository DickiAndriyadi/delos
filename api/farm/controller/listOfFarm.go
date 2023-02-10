package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/response"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *FarmController) ListOfFarm(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

	data, err := ctrl.service.ListOfFarm()
	if err != nil {
		return echo.NewHTTPError(
			err.ErrCode,
			err.Message.Error(),
		)
	}

	result.Data = map[string]interface{}{
		"items": data,
	}

	return c.JSON(http.StatusOK, result)
}
