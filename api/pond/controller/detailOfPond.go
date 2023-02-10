package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/response"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *PondController) DetailOfPond(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

	id := c.Param("id")

	data, err := ctrl.service.DetailOfPond(id)
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
