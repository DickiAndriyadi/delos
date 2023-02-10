package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/payload"
	"delos/api/model/response"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *FarmController) UpdateFarm(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

	id := c.Param("id")

	payload := new(payload.AddUpdateFarm)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := payload.Validate(); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err,
		)
	}

	data, err := ctrl.service.UpdateFarm(*payload, id)
	if err != nil {
		return echo.NewHTTPError(
			err.ErrCode,
			err.Message.Error(),
		)
	}

	result.Data = map[string]interface{}{
		"message": data,
	}

	return c.JSON(http.StatusOK, result)
}
