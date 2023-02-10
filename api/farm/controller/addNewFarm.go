package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/payload"
	"delos/api/model/response"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *FarmController) AddNewFarm(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

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

	data, err := ctrl.service.AddNewFarm(*payload)
	if err != nil {
		return echo.NewHTTPError(
			err.ErrCode,
			err.Message.Error(),
		)
	}

	result.Data = data

	return c.JSON(http.StatusCreated, result)
}
